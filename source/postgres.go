package source

import (
	"context"
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/summit-fi/wordsdk-go/utils/locale"

	"io"
)

type Postgres struct {
	ctx  context.Context
	pool *pgxpool.Pool
}

func NewPostgres(ctx context.Context, connString string, loadRemoteFiles []string) (*Postgres, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}

	conn := Postgres{
		ctx:  ctx,
		pool: pool,
	}

	if len(loadRemoteFiles) > 0 {

		var openedFiles []openedFile

		for _, filePath := range loadRemoteFiles {
			f, err := os.Open(filePath)
			if err != nil {
				return nil, fmt.Errorf("failed to open remote file '%s': %v", filePath, err)
			}
			defer f.Close()

			read, err := io.ReadAll(f)
			if err != nil {
				return nil, fmt.Errorf("failed to read remote file '%s': %v", filePath, err)
			}

			openedFiles = append(openedFiles, openedFile{
				localeCode: locale.GetLocaleFromFileName(f.Name()),
				bytes:      read,
			})

		}

		for _, f := range openedFiles {

			parsed := FtlParse(f.localeCode, f.bytes)

			err = conn.batchSaveTranslations(parsed)
			if err != nil {
				return nil, fmt.Errorf("failed to save translations from remote files: %v", err)
			}
		}

	}
	return &conn, nil
}

func computeChecksum(translations []Object) string {
	h := sha256.New()

	for _, t := range translations {
		io.WriteString(h, t.LocaleCode)
		io.WriteString(h, t.Key)
		io.WriteString(h, t.Value)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

func (p *Postgres) LoadAllStatic(checksumIn string) (result []Object, checksumOut string, err error) {

	object, err := p.getAllKeys()
	if err != nil {
		return nil, "", err
	}
	checksumOut = computeChecksum(object)

	if checksumIn == checksumOut {
		return result, checksumOut, nil
	}

	return result, checksumOut, nil
}

func (p *Postgres) LoadAllDynamic(accessKey, checksumIn string) (result []Object, checkSumOut string, err error) {
	return p.LoadAllStatic(checksumIn)
}

func (p *Postgres) LoadOneDynamic(accessKey, lang, key string) (string, error) {
	value, err := p.getTranslation(lang, key)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (p *Postgres) SaveDynamic(accessKey string, data []Object) error {
	var dataMap = make(map[string]map[string]interface{}) // localeCode -> key -> value
	for _, datum := range data {
		if _, ok := dataMap[datum.LocaleCode]; !ok {
			dataMap[datum.LocaleCode] = make(map[string]interface{})
		}
		dataMap[datum.LocaleCode][datum.Key] = datum.Value
	}

	var translations []Object
	for localeCode, keyValues := range dataMap {
		for key, value := range keyValues {
			if strValue, ok := value.(string); ok {
				translations = append(translations, Object{
					LocaleCode: localeCode,
					Key:        key,
					Value:      strValue,
				})
			}
		}
	}

	return p.batchSaveTranslations(translations)
}

func (p *Postgres) Close() {
	p.pool.Close()
}

func (p *Postgres) getAllKeys() ([]Object, error) {
	rows, err := p.pool.Query(p.ctx, "SELECT lang, code, value FROM translation;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var datum []Object
	for rows.Next() {
		var obj Object
		if err := rows.Scan(
			&obj.LocaleCode,
			&obj.Key,
			&obj.Value); err != nil {
			return nil, err
		}
		datum = append(datum, obj)
	}
	return datum, nil
}

func (p *Postgres) getTranslation(lang, key string) (string, error) {
	var value string
	err := p.pool.QueryRow(p.ctx, "SELECT value FROM translation WHERE lang = $1 AND code = $2", lang, key).Scan(&value)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (p *Postgres) saveTranslation(ctx context.Context, lang, key, value string) error {
	keyType := "content"
	_, err := p.pool.Exec(ctx, "INSERT INTO translation (type, lang, code, value) VALUES ($1, $2, $3, $4) ON CONFLICT (lang, code) DO UPDATE SET value = EXCLUDED.value",
		keyType,
		lang,
		key,
		value,
	)
	return err
}

func (p *Postgres) batchSaveTranslations(translations []Object) error {
	tx, err := p.pool.Begin(p.ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(p.ctx)

	keyType := "content"
	for _, t := range translations {
		_, err := tx.Exec(p.ctx, "INSERT INTO translation (type, lang, code, value) VALUES ($1, $2, $3, $4) ON CONFLICT (lang, code) DO UPDATE SET value = EXCLUDED.value",
			keyType,
			t.LocaleCode,
			t.Key,
			t.Value,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit(p.ctx)
}
