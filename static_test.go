package word

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/summit-fi/wordsdk-go/source"
)

func ftlClientWithSaveStrategy(saveStrategy SaveStrategy) (SDK, error) {

	db := source.NewFtl()
	err := db.AddLocaleFiles(map[string]string{
		"uk_UA": filepath.Join(Root(), "test", "fixtures", "custom", "custom_data", "uk_UA.ftl"),
		"en_UA": filepath.Join(Root(), "test", "fixtures", "custom", "custom_data", "en_UA.ftl"),
		"es_CO": filepath.Join(Root(), "test", "fixtures", "custom", "custom_data", "es_CO.ftl"),
		"en_EU": filepath.Join(Root(), "test", "fixtures", "custom", "custom_data", "en_EU.ftl"),
	})
	if err != nil {
		return nil, err
	}

	c := Config{
		Source:       db,
		SaveStrategy: saveStrategy,
	}
	return NewClient(&c)
}

func TestClient_Flush(t *testing.T) {

	// Init client
	w, err := ftlClientWithSaveStrategy(SaveStrategyOnDemand)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	// Read initial file to compare
	f, err := os.Open(filepath.Join(Root(), "test", "fixtures", "custom", "custom_data", "en_EU.ftl"))
	if err != nil {
		t.Fatalf("os.Open() error = %v", err)
	}

	// Read initial file
	b, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("io.ReadAll() error = %v", err)
	}
	f.Close()
	initialData := string(b)

	// Save new data
	err = w.SaveTranslations([]source.Object{
		{
			LocaleCode: "en_US",
			Key:        "core.every",
			Value:      "Every 1 minute",
		},
	})

	// Read file after save
	f, err = os.Open(filepath.Join(Root(), "test", "fixtures", "custom", "custom_data", "en_EU.ftl"))
	if err != nil {
		t.Fatalf("os.Open() error = %v", err)
	}

	// Read file after save
	b, err = io.ReadAll(f)
	if err != nil {
		t.Fatalf("io.ReadAll() error = %v", err)
	}
	f.Close()
	updatedData := string(b)

	// Data should be updated after Flush
	updatedDataShouldBe := fmt.Sprintf("%s\n%s = %s\n", initialData, "core.every", "Every 2 minute")
	if updatedData != updatedDataShouldBe {
		t.Fatalf("updatedData = %v \n\n\n\n want %v", updatedData, updatedDataShouldBe)
	}

	// Flush
	err = w.Flush()
	if err != nil {
		t.Fatalf("Flush() error = %v", err)
	}

	// Read file after Flush
	f, err = os.Open(filepath.Join(Root(), "test", "fixtures", "custom", "custom_data", "en_EU.ftl"))
	if err != nil {
		t.Fatalf("os.Open() error = %v", err)
	}

	// Read file after Flush
	b, err = io.ReadAll(f)
	if err != nil {
		t.Fatalf("io.ReadAll() error = %v", err)
	}
	f.Close()
	updatedData = string(b)

	if updatedData != updatedDataShouldBe {
		t.Fatalf("updatedData = %v \n\n\n\n want %v", updatedData, updatedDataShouldBe)
	}

	// Write initial file back
	f, err = os.OpenFile(filepath.Join(Root(), "test", "fixtures", "custom", "custom_data", "en_EU.ftl"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		t.Fatalf("os.OpenFile() error = %v", err)
	}

	_, err = f.WriteString(initialData)
	if err != nil {
		t.Fatalf("f.WriteString() error = %v", err)
	}

	f.Close()
}

func TestClient_SaveTranslation(t *testing.T) {

	// Init client
	w, err := ftlClientWithSaveStrategy(SaveStrategyImmediate)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	// Read initial file to compare
	f, err := os.Open(filepath.Join(Root(), "test", "fixtures", "custom", "custom_data", "en_EU.ftl"))
	if err != nil {
		t.Fatalf("os.Open() error = %v", err)
	}

	// Read initial file
	b, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("io.ReadAll() error = %v", err)
	}
	f.Close()
	initialData := string(b)

	// Save new data
	err = w.SaveTranslations([]source.Object{
		{
			LocaleCode: "en_EU",
			Key:        "core.every",
			Value:      "Every 2 minute",
		},
	})

	// Read file after save
	f, err = os.Open(filepath.Join(Root(), "test", "fixtures", "custom", "custom_data", "en_EU.ftl"))
	if err != nil {
		t.Fatalf("os.Open() error = %v", err)
	}

	// Read file after save
	b, err = io.ReadAll(f)
	if err != nil {
		t.Fatalf("io.ReadAll() error = %v", err)
	}
	f.Close()
	updatedData := string(b)

	// Data should be updated after Flush
	updatedDataShouldBe := fmt.Sprintf("%s\n%s = %s\n", initialData, "core.every", "Every 2 minute")

	if updatedData != updatedDataShouldBe {
		t.Fatalf("updatedData = %v \n\n\n\n want %v", updatedData, updatedDataShouldBe)
	}

	// Write initial file back
	f, err = os.OpenFile(filepath.Join(Root(), "test", "fixtures", "custom", "custom_data", "en_EU.ftl"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		t.Fatalf("os.OpenFile() error = %v", err)
	}

	_, err = f.WriteString(initialData)
	if err != nil {
		t.Fatalf("f.WriteString() error = %v", err)
	}

	f.Close()
}

func TestClient_Reset(t *testing.T) {
	connect, err := ftlClientWithSaveStrategy(SaveStrategyOnDemand)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	connect = connect.EnableDynamicContent(XKeyGen("S", "summit", "schedule"))
	t.Logf("X-Dynamic-Key: %s", XKeyGen("S", "summit", "schedule"))

	err = connect.Dynamic().SaveTranslation("en_EU", "test_reset", "Testing reset")
	if err != nil {
		t.Errorf("Failed to save translation: %v", err)
		return
	}

	connect.Reset()

	str := connect.Dynamic().T("en_EU", "test_reset")
	if str != "test_reset" {
		t.Errorf("Failed to reset translation: %v", str)
		return
	}

	t.Logf("Translation after reset: %s", str)
}
