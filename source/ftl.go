package source

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/summit-fi/wordsdk-go/fluent/parser"
	"github.com/summit-fi/wordsdk-go/fluent/parser/ast"
)

type Ftl struct {
	files []file
	sync.RWMutex
}

const ftlBlockIndent = "    "

const ftlEmptyValue = `{ "" }`

// FormatFTLEntry formats one key/value pair as a valid FTL entry.
// Multi-line values are emitted as block patterns so line breaks are part of
// the value instead of being interpreted as the end of the entry.
func FormatFTLEntry(key, value string) string {
	value = normalizeFTLNewlines(value)
	if value == "" {
		value = ftlEmptyValue
	}
	if !strings.Contains(value, "\n") {
		return fmt.Sprintf("%s = %s\n", key, value)
	}

	var builder strings.Builder
	builder.WriteString(key)
	builder.WriteString(" =\n")
	for _, line := range strings.Split(value, "\n") {
		builder.WriteString(ftlBlockIndent)
		builder.WriteString(line)
		builder.WriteString("\n")
	}
	return builder.String()
}

func normalizeFTLNewlines(value string) string {
	value = strings.ReplaceAll(value, "\r\n", "\n")
	return strings.ReplaceAll(value, "\r", "\n")
}

func (f *Ftl) SaveDynamic(_ string, data []Object) error {
	f.Lock()
	defer f.Unlock()

	updatesByLocale := groupFTLUpdates(data)
	type pendingWrite struct {
		path string
		data []byte
	}
	pending := make([]pendingWrite, 0, len(f.files))

	for _, fi := range f.files {
		updates := updatesByLocale[fi.localeCode]
		if len(updates) == 0 {
			// File with this locale is not loaded/attached
			continue
		}

		content, err := os.ReadFile(fi.path)
		if err != nil {
			return fmt.Errorf("read FTL file %s: %w", fi.path, err)
		}

		updated, err := updateFTLSource(content, updates)
		if err != nil {
			return fmt.Errorf("update FTL file %s: %w", fi.path, err)
		}

		pending = append(pending, pendingWrite{path: fi.path, data: updated})
	}

	for _, write := range pending {
		if err := writeFTLFile(write.path, write.data); err != nil {
			return fmt.Errorf("write FTL file %s: %w", write.path, err)
		}
	}

	return nil
}

func (f *Ftl) LoadAllDynamic(key string, checksumIn string) (result []Object, checksumOut string, err error) {
	return f.LoadAllStatic(checksumIn)
}
func (f *Ftl) LoadOneDynamic(accessKey, lang, key string) (string, error) {
	f.RLock()
	defer f.RUnlock()

	for _, file := range f.files {
		if file.localeCode == lang {
			data, err := os.ReadFile(file.path)
			if err != nil {
				return key, fmt.Errorf("read FTL file %s: %w", file.path, err)
			}

			entries, err := parseFTLEntries(lang, data)
			if err != nil {
				return key, fmt.Errorf("parse FTL file %s: %w", file.path, err)
			}
			for _, entry := range entries {
				if entry.object.Key == key {
					return entry.patternValue, nil
				}
			}
		}
	}

	return key, fmt.Errorf("key %q not found for locale %q", key, lang)
}

type file struct {
	localeCode string
	path       string
}

func NewFtl() *Ftl {
	return &Ftl{}
}

func (f *Ftl) AddLocaleFile(localeCode string, path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("stat FTL file %s: %w", path, err)
	}
	if info.IsDir() {
		return fmt.Errorf("FTL path %s is a directory", path)
	}

	f.Lock()
	defer f.Unlock()
	for _, existing := range f.files {
		existingInfo, statErr := os.Stat(existing.path)
		if statErr != nil {
			return fmt.Errorf("stat attached FTL file %s: %w", existing.path, statErr)
		}
		if os.SameFile(info, existingInfo) {
			if existing.localeCode == localeCode {
				return nil
			}
			return fmt.Errorf(
				"FTL file %s is already attached to locale %s",
				existing.path,
				existing.localeCode,
			)
		}
	}
	f.files = append(f.files, file{
		localeCode: localeCode,
		path:       path,
	})
	return nil
}

func (f *Ftl) AddLocaleFiles(paths map[string]string) error {
	locales := make([]string, 0, len(paths))
	for locale := range paths {
		locales = append(locales, locale)
	}
	sort.Strings(locales)

	for _, locale := range locales {
		path := paths[locale]
		err := f.AddLocaleFile(locale, path)
		if err != nil {
			return err
		}
	}
	return nil
}

type openedFile struct {
	localeCode string
	bytes      []byte
}

type parsedFTLEntry struct {
	object       Object
	patternValue string
}

// FtlParse is the compatibility wrapper for callers that cannot handle parse
// errors.
// Deprecated: use FtlParseWithError.
func FtlParse(locale string, data []byte) []Object {
	objects, _ := FtlParseWithError(locale, data)
	return objects
}

// FtlParseWithError parses FTL with the repository's Fluent parser. Values are
// serialized as valid right-hand sides so they can be fed back into a bundle.
func FtlParseWithError(locale string, data []byte) ([]Object, error) {
	entries, err := parseFTLEntries(locale, data)
	objects := make([]Object, 0, len(entries))
	for _, entry := range entries {
		objects = append(objects, entry.object)
	}
	return objects, err
}

func parseFTLEntries(locale string, data []byte) ([]parsedFTLEntry, error) {
	if !utf8.Valid(data) {
		return nil, fmt.Errorf("FTL source is not valid UTF-8")
	}
	source := normalizeFTLNewlines(string(data))
	sourceRunes := []rune(source)
	resource, parseErrors := parseFTLResource(source)
	duplicateError := validateUniqueFTLEntries(resource)

	entries := make([]parsedFTLEntry, 0, len(resource.Body))
	entryIndexes := make(map[string]int)
	for _, node := range resource.Body {
		var (
			key        string
			value      *ast.Pattern
			attributes []*ast.Attribute
		)

		switch entry := node.(type) {
		case *ast.Message:
			key = entry.ID.Name
			value = entry.Value
			attributes = entry.Attributes
		case *ast.Term:
			key = "-" + entry.ID.Name
			value = entry.Value
			attributes = entry.Attributes
		default:
			continue
		}

		serialized, patternValue, err := serializeFTLEntryValue(sourceRunes, value, attributes)
		if err != nil {
			return nil, err
		}
		dynamicValue := patternValue
		if isEmptyFTLPattern(value) {
			dynamicValue = ""
		}
		parsed := parsedFTLEntry{
			object: Object{
				LocaleCode: locale,
				Key:        key,
				Value:      serialized,
			},
			patternValue: dynamicValue,
		}

		// Preserve source order while retaining the old last-definition-wins
		// behavior for duplicate keys.
		if index, exists := entryIndexes[key]; exists {
			entries[index] = parsed
			continue
		}
		entryIndexes[key] = len(entries)
		entries = append(entries, parsed)
	}

	if len(parseErrors) > 0 {
		return entries, formatFTLParseError(sourceRunes, parseErrors)
	}
	if duplicateError != nil {
		return entries, duplicateError
	}
	return entries, nil
}

func serializeFTLEntryValue(source []rune, value *ast.Pattern, attributes []*ast.Attribute) (string, string, error) {
	var builder strings.Builder
	patternValue := ""
	if value != nil {
		var err error
		patternValue, err = serializeFTLPattern(source, value)
		if err != nil {
			return "", "", err
		}
		builder.WriteString(formatFTLRightHandPattern(patternValue))
	}

	for _, attribute := range attributes {
		attributeValue, err := serializeFTLPattern(source, attribute.Value)
		if err != nil {
			return "", "", err
		}
		if attributeValue == "" {
			attributeValue = ftlEmptyValue
		}

		builder.WriteString("\n")
		builder.WriteString(ftlBlockIndent)
		builder.WriteString(".")
		builder.WriteString(attribute.ID.Name)
		if strings.Contains(attributeValue, "\n") {
			builder.WriteString(" =\n")
			builder.WriteString(indentFTLLines(attributeValue, ftlBlockIndent+ftlBlockIndent))
		} else {
			builder.WriteString(" = ")
			builder.WriteString(attributeValue)
		}
	}

	return builder.String(), patternValue, nil
}

func serializeFTLPattern(source []rune, pattern *ast.Pattern) (string, error) {
	var builder strings.Builder
	for _, element := range pattern.Elements {
		switch element := element.(type) {
		case *ast.Text:
			builder.WriteString(element.Value)
		case *ast.Placeable:
			start := int(element.Span[0])
			end := int(element.Span[1])
			if start < 0 || end < start || end > len(source) {
				return "", fmt.Errorf("invalid placeable span [%d:%d]", start, end)
			}
			builder.WriteString(string(source[start:end]))
		default:
			return "", fmt.Errorf("unsupported FTL pattern element %T", element)
		}
	}
	return builder.String(), nil
}

func isEmptyFTLPattern(pattern *ast.Pattern) bool {
	if pattern == nil || len(pattern.Elements) != 1 {
		return false
	}
	placeable, ok := pattern.Elements[0].(*ast.Placeable)
	if !ok {
		return false
	}
	literal, ok := placeable.Expression.(*ast.StringLiteral)
	return ok && literal.Value == ""
}

func formatFTLRightHandPattern(value string) string {
	if value == "" {
		return ftlEmptyValue
	}
	if !strings.Contains(value, "\n") {
		return value
	}
	return "\n" + indentFTLLines(value, ftlBlockIndent)
}

func indentFTLLines(value, indent string) string {
	lines := strings.Split(normalizeFTLNewlines(value), "\n")
	var builder strings.Builder
	for index, line := range lines {
		if index > 0 {
			builder.WriteString("\n")
		}
		builder.WriteString(indent)
		builder.WriteString(line)
	}
	return builder.String()
}

func formatFTLParseError(source []rune, parseErrors []*parser.Error) error {
	first := parseErrors[0]
	line, column := 1, 1
	limit := min(int(first.Span[0]), len(source))
	for _, char := range source[:limit] {
		if char == '\n' {
			line++
			column = 1
			continue
		}
		column++
	}

	if len(parseErrors) == 1 {
		return fmt.Errorf("invalid FTL at line %d, column %d: %s", line, column, first.Message)
	}
	return fmt.Errorf(
		"invalid FTL at line %d, column %d: %s (and %d more errors)",
		line,
		column,
		first.Message,
		len(parseErrors)-1,
	)
}

func parseFTLResource(source string) (*ast.Resource, []*parser.Error) {
	return parser.New(maskTopLevelFTLComments(source)).Parse()
}

// The previous loader accepted comments without the Fluent-required space
// (for example, "#comment"). Mask comment contents for parsing while keeping
// every rune position unchanged, so existing files remain loadable and AST
// spans still point into the original source.
func maskTopLevelFTLComments(source string) string {
	runes := []rune(source)
	lineStart := true
	for index := 0; index < len(runes); index++ {
		if lineStart && runes[index] == '#' {
			for index++; index < len(runes) && runes[index] != '\n'; index++ {
				runes[index] = ' '
			}
		}
		lineStart = index < len(runes) && runes[index] == '\n'
	}
	return string(runes)
}

func groupFTLUpdates(data []Object) map[string][]Object {
	grouped := make(map[string][]Object)
	indexes := make(map[string]map[string]int)
	for _, object := range data {
		localeIndexes := indexes[object.LocaleCode]
		if localeIndexes == nil {
			localeIndexes = make(map[string]int)
			indexes[object.LocaleCode] = localeIndexes
		}

		if index, exists := localeIndexes[object.Key]; exists {
			grouped[object.LocaleCode][index] = object
			continue
		}
		localeIndexes[object.Key] = len(grouped[object.LocaleCode])
		grouped[object.LocaleCode] = append(grouped[object.LocaleCode], object)
	}
	return grouped
}

type ftlReplacement struct {
	start int
	end   int
	text  string
}

type validatedFTLUpdate struct {
	object         Object
	formattedEntry string
	hasAttributes  bool
}

func updateFTLSource(data []byte, updates []Object) ([]byte, error) {
	if !utf8.Valid(data) {
		return nil, fmt.Errorf("FTL source is not valid UTF-8")
	}
	original := string(data)
	lineEnding := preferredFTLLineEnding(original)
	source := []rune(normalizeFTLNewlines(original))
	resource, parseErrors := parseFTLResource(string(source))
	if len(parseErrors) > 0 {
		return nil, formatFTLParseError(source, parseErrors)
	}
	if err := validateUniqueFTLEntries(resource); err != nil {
		return nil, err
	}

	updatesByKey := make(map[string]validatedFTLUpdate, len(updates))
	for _, update := range updates {
		validated, err := validateFTLUpdate(update)
		if err != nil {
			return nil, err
		}
		updatesByKey[update.Key] = validated
	}

	matched := make(map[string]bool, len(updates))
	replacements := make([]ftlReplacement, 0, len(updates))
	for _, node := range resource.Body {
		var (
			key        string
			identifier *ast.Identifier
			value      *ast.Pattern
			attributes []*ast.Attribute
			entryStart int
			entryEnd   int
		)

		switch entry := node.(type) {
		case *ast.Message:
			key = entry.ID.Name
			identifier = entry.ID
			value = entry.Value
			attributes = entry.Attributes
			entryStart = int(entry.ID.Span[0])
			entryEnd = int(entry.Span[1])
		case *ast.Term:
			key = "-" + entry.ID.Name
			identifier = entry.ID
			value = entry.Value
			attributes = entry.Attributes
			entryStart = int(entry.ID.Span[0]) - 1
			entryEnd = int(entry.Span[1])
		default:
			continue
		}

		update, exists := updatesByKey[key]
		if !exists {
			continue
		}
		var replacement ftlReplacement
		var err error
		if update.hasAttributes {
			replacement, err = replaceWholeFTLEntry(source, entryStart, entryEnd, update.formattedEntry)
		} else {
			replacement, err = replaceFTLEntryValue(source, identifier, value, attributes, update.object.Value)
		}
		if err != nil {
			return nil, fmt.Errorf("update key %q: %w", key, err)
		}
		replacements = append(replacements, replacement)
		matched[key] = true
	}

	sort.Slice(replacements, func(i, j int) bool {
		return replacements[i].start < replacements[j].start
	})
	var builder strings.Builder
	cursor := 0
	for _, replacement := range replacements {
		if replacement.start < cursor || replacement.end < replacement.start || replacement.end > len(source) {
			return nil, fmt.Errorf("overlapping or invalid FTL replacement [%d:%d]", replacement.start, replacement.end)
		}
		builder.WriteString(string(source[cursor:replacement.start]))
		builder.WriteString(replacement.text)
		cursor = replacement.end
	}
	builder.WriteString(string(source[cursor:]))
	updated := builder.String()

	for _, update := range updates {
		if matched[update.Key] {
			continue
		}
		if updated != "" {
			updated += "\n"
		}
		updated += FormatFTLEntry(update.Key, update.Value)
	}

	_, finalErrors := parseFTLResource(updated)
	if len(finalErrors) > 0 {
		return nil, fmt.Errorf("updated content is invalid: %w", formatFTLParseError([]rune(updated), finalErrors))
	}
	if lineEnding == "\r\n" {
		updated = strings.ReplaceAll(updated, "\n", "\r\n")
	}
	return []byte(updated), nil
}

func replaceFTLEntryValue(
	source []rune,
	identifier *ast.Identifier,
	value *ast.Pattern,
	attributes []*ast.Attribute,
	newValue string,
) (ftlReplacement, error) {
	newValue = normalizeFTLNewlines(newValue)
	if newValue == "" {
		newValue = ftlEmptyValue
	}

	if value != nil {
		start := int(value.Span[0])
		end := int(value.Span[1])
		if start < 0 || end < start || end > len(source) {
			return ftlReplacement{}, fmt.Errorf("invalid pattern span [%d:%d]", start, end)
		}
		blockPattern := start > 0 && source[start-1] == '\n'
		return ftlReplacement{
			start: start,
			end:   end,
			text:  formatFTLPatternForSpan(newValue, blockPattern),
		}, nil
	}

	if len(attributes) == 0 {
		return ftlReplacement{}, fmt.Errorf("entry has neither a value nor attributes")
	}
	start := int(identifier.Span[1])
	end := int(attributes[0].Span[0])
	if start < 0 || end < start || end > len(source) {
		return ftlReplacement{}, fmt.Errorf("invalid attribute-only entry span [%d:%d]", start, end)
	}

	equals := -1
	for index := start; index < end; index++ {
		if source[index] == '=' {
			equals = index
			break
		}
	}
	if equals < 0 {
		return ftlReplacement{}, fmt.Errorf("entry assignment was not found")
	}

	return ftlReplacement{
		start: equals + 1,
		end:   end,
		text:  formatFTLValueBeforeAttribute(newValue),
	}, nil
}

func replaceWholeFTLEntry(
	source []rune,
	entryStart int,
	entryEnd int,
	formattedEntry string,
) (ftlReplacement, error) {
	if entryStart < 0 || entryEnd < entryStart || entryEnd > len(source) {
		return ftlReplacement{}, fmt.Errorf("invalid entry span [%d:%d]", entryStart, entryEnd)
	}
	return ftlReplacement{
		start: entryStart,
		end:   entryEnd,
		text:  strings.TrimSuffix(formattedEntry, "\n"),
	}, nil
}

func formatFTLPatternForSpan(value string, block bool) string {
	if block {
		return indentFTLLines(value, ftlBlockIndent)
	}
	lines := strings.Split(value, "\n")
	if len(lines) == 1 {
		return value
	}
	return lines[0] + "\n" + indentFTLLines(strings.Join(lines[1:], "\n"), ftlBlockIndent)
}

func formatFTLValueBeforeAttribute(value string) string {
	if strings.Contains(value, "\n") {
		return "\n" + indentFTLLines(value, ftlBlockIndent) + "\n" + ftlBlockIndent
	}
	return " " + value + "\n" + ftlBlockIndent
}

func validateFTLUpdate(update Object) (validatedFTLUpdate, error) {
	if !utf8.ValidString(update.Key) || !utf8.ValidString(update.Value) {
		return validatedFTLUpdate{}, fmt.Errorf("update for key %q is not valid UTF-8", update.Key)
	}
	formatted := FormatFTLEntry(update.Key, update.Value)
	resource, parseErrors := parseFTLResource(formatted)
	if len(parseErrors) > 0 {
		return validatedFTLUpdate{}, fmt.Errorf(
			"invalid update for key %q: %w",
			update.Key,
			formatFTLParseError([]rune(formatted), parseErrors),
		)
	}

	var validated *validatedFTLUpdate
	for _, node := range resource.Body {
		switch entry := node.(type) {
		case *ast.Message:
			if update.Key == entry.ID.Name && (entry.Value != nil || len(entry.Attributes) > 0) {
				candidate := validatedFTLUpdate{
					object:         update,
					formattedEntry: formatted,
					hasAttributes:  len(entry.Attributes) > 0,
				}
				validated = &candidate
			}
		case *ast.Term:
			if update.Key == "-"+entry.ID.Name && (entry.Value != nil || len(entry.Attributes) > 0) {
				candidate := validatedFTLUpdate{
					object:         update,
					formattedEntry: formatted,
					hasAttributes:  len(entry.Attributes) > 0,
				}
				validated = &candidate
			}
		}
	}
	if validated == nil || len(resource.Body) != 1 {
		return validatedFTLUpdate{}, fmt.Errorf("invalid FTL key or value for %q", update.Key)
	}
	return *validated, nil
}

func preferredFTLLineEnding(value string) string {
	if strings.Contains(value, "\r\n") && !strings.Contains(strings.ReplaceAll(value, "\r\n", ""), "\n") {
		return "\r\n"
	}
	return "\n"
}

func validateUniqueFTLEntries(resource *ast.Resource) error {
	seen := make(map[string]struct{})
	for _, node := range resource.Body {
		var key string
		switch entry := node.(type) {
		case *ast.Message:
			key = entry.ID.Name
		case *ast.Term:
			key = "-" + entry.ID.Name
		default:
			continue
		}
		if _, exists := seen[key]; exists {
			return fmt.Errorf("duplicate FTL key %q", key)
		}
		seen[key] = struct{}{}
	}
	return nil
}

// Resolve symlinks before the atomic rename so the registered symlink remains
// intact and readers never observe a partially written file.
func writeFTLFile(path string, data []byte) error {
	target, err := filepath.EvalSymlinks(path)
	if err != nil {
		return err
	}
	info, err := os.Stat(target)
	if err != nil {
		return err
	}
	temporary, err := os.CreateTemp(filepath.Dir(target), "."+filepath.Base(target)+".tmp-*")
	if err != nil {
		return err
	}
	temporaryPath := temporary.Name()
	defer os.Remove(temporaryPath)

	if err := temporary.Chmod(info.Mode()); err != nil {
		temporary.Close()
		return err
	}
	if _, err := temporary.Write(data); err != nil {
		temporary.Close()
		return err
	}
	if err := temporary.Sync(); err != nil {
		temporary.Close()
		return err
	}
	if err := temporary.Close(); err != nil {
		return err
	}
	return os.Rename(temporaryPath, target)
}

func (f *Ftl) LoadAllStatic(checksumIn string) ([]Object, string, error) {
	f.RLock()
	defer f.RUnlock()

	files := make([]openedFile, 0, len(f.files))
	h := sha256.New()
	var length [8]byte
	for _, attached := range f.files {
		content, err := os.ReadFile(attached.path)
		if err != nil {
			return nil, "", fmt.Errorf("read FTL file %s: %w", attached.path, err)
		}

		binary.LittleEndian.PutUint64(length[:], uint64(len(attached.localeCode)))
		h.Write(length[:])
		h.Write([]byte(attached.localeCode))
		binary.LittleEndian.PutUint64(length[:], uint64(len(content)))
		h.Write(length[:])
		h.Write(content)
		files = append(files, openedFile{
			localeCode: attached.localeCode,
			bytes:      content,
		})
	}

	checksum := fmt.Sprintf("%x", h.Sum(nil))
	if checksum == checksumIn {
		return nil, checksum, nil
	}

	var objects []Object
	for _, file := range files {
		parsed, err := FtlParseWithError(file.localeCode, file.bytes)
		if err != nil {
			return nil, "", fmt.Errorf("parse locale %s: %w", file.localeCode, err)
		}
		objects = append(objects, parsed...)
	}

	return objects, checksum, nil
}
