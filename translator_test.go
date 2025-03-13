package word

import (
	"io"
	"os"
	"testing"

	"github.com/summit-fi/wordsdk-go/source"
)

func jsonClientWithSaveStrategy(saveStrategy SaveStrategy) (SDK, error) {

	db := source.NewJson()
	err := db.AddLocaleFiles(map[string]string{
		"uk_UA": "example/assets/json/uk.json",
		"en_US": "example/assets/json/en.json",
		"es_ES": "example/assets/json/es.json",
		"sv_SE": "example/assets/json/sv.json",
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
	w, err := jsonClientWithSaveStrategy(SaveStrategyOnDemand)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	// Read initial file to compare
	f, err := os.Open("example/assets/json/en.json")
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

	initialDataShouldBe := `{
  "core.every": "Every",
  "core.everyday": "Every day",
  "notifications.email.confirmation.title": "Confirmation"
}`
	if initialData != initialDataShouldBe {
		t.Fatalf("initialData = %v, want %v", initialData, initialDataShouldBe)
	}

	// Save new data
	err = w.SaveTranslations([]source.Object{
		{
			LocaleCode: "en_US",
			Key:        "core.every",
			Value:      "Every 1 minute",
		},
	})

	// Read file after save
	f, err = os.Open("example/assets/json/en.json")
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

	// Data should not be updated before Flush
	if updatedData != initialData {
		t.Fatalf("updatedData = %v, want %v", updatedData, initialData)
	}

	// Flush
	err = w.Flush()
	if err != nil {
		t.Fatalf("Flush() error = %v", err)
	}

	// Read file after Flush
	f, err = os.Open("example/assets/json/en.json")
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

	// Data should be updated after Flush
	updatedDataShouldBe := `{
  "core.every": "Every 1 minute",
  "core.everyday": "Every day",
  "notifications.email.confirmation.title": "Confirmation"
}`

	if updatedData != updatedDataShouldBe {
		t.Fatalf("updatedData = %v, want %v", updatedData, updatedDataShouldBe)
	}

	// Write initial file back
	f, err = os.OpenFile("example/assets/json/en.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
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
	w, err := jsonClientWithSaveStrategy(SaveStrategyImmediate)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	// Read initial file to compare
	f, err := os.Open("example/assets/json/en.json")
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

	initialDataShouldBe := `{
  "core.every": "Every",
  "core.everyday": "Every day",
  "notifications.email.confirmation.title": "Confirmation"
}`
	if initialData != initialDataShouldBe {
		t.Fatalf("initialData = %v, want %v", initialData, initialDataShouldBe)
	}

	// Save new data
	err = w.SaveTranslations([]source.Object{
		{
			LocaleCode: "en_US",
			Key:        "core.every",
			Value:      "Every 2 minute",
		},
	})

	// Read file after save
	f, err = os.Open("example/assets/json/en.json")
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
	updatedDataShouldBe := `{
  "core.every": "Every 2 minute",
  "core.everyday": "Every day",
  "notifications.email.confirmation.title": "Confirmation"
}`

	if updatedData != updatedDataShouldBe {
		t.Fatalf("updatedData = %v, want %v", updatedData, updatedDataShouldBe)
	}

	// Write initial file back
	f, err = os.OpenFile("example/assets/json/en.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		t.Fatalf("os.OpenFile() error = %v", err)
	}

	_, err = f.WriteString(initialData)
	if err != nil {
		t.Fatalf("f.WriteString() error = %v", err)
	}

	f.Close()
}
