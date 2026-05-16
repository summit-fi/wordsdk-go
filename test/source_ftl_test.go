package test

import (
	"os"
	"strings"
	"testing"

	"github.com/summit-fi/wordsdk-go/source"
)

func createTempFile(t *testing.T, content string) (string, func()) {
	tmp, err := os.CreateTemp("", "example_ftl_*.ftl")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tmp.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmp.Close()
	return tmp.Name(), func() { os.Remove(tmp.Name()) }
}

func TestSaveDynamicAddsNewKeyValuePairs(t *testing.T) {
	path, cleanup := createTempFile(t, "")
	defer cleanup()

	ftl := source.NewFtl()
	_ = ftl.AddLocaleFile("en_EU", path)

	data := []source.Object{
		{LocaleCode: "en_EU", Key: "greet", Value: "hello"},
		{LocaleCode: "en_EU", Key: "farewell", Value: "goodbye"},
	}
	err := ftl.SaveDynamic("", data)
	if err != nil {
		t.Fatalf("SaveDynamic error: %v", err)
	}

	val, err := ftl.LoadOneDynamic("", "en_EU", "greet")
	if err != nil {
		t.Fatalf("LoadOneDynamic error: %v", err)
	}
	if val != "hello" {
		t.Errorf("expected 'hello', got '%s'", val)
	}

	val, err = ftl.LoadOneDynamic("", "en_EU", "farewell")
	if err != nil {
		t.Fatalf("LoadOneDynamic error: %v", err)
	}
	if val != "goodbye" {
		t.Errorf("expected 'goodbye', got '%s'", val)
	}
}

func TestSaveDynamicIgnoresUnloadedLocales(t *testing.T) {
	path, cleanup := createTempFile(t, "")
	defer cleanup()

	ftl := source.NewFtl()
	_ = ftl.AddLocaleFile("en_EU", path)

	data := []source.Object{
		{LocaleCode: "fr", Key: "greet", Value: "bonjour"},
	}
	err := ftl.SaveDynamic("", data)
	if err != nil {
		t.Fatalf("SaveDynamic error: %v", err)
	}

	val, err := ftl.LoadOneDynamic("", "fr", "greet")
	if err == nil {
		t.Fatalf("expected error for unloaded locale, got value: %s", val)
	}
}

func TestLoadOneDynamicReturnsErrorForMissingKey(t *testing.T) {
	path, cleanup := createTempFile(t, "")
	defer cleanup()

	ftl := source.NewFtl()
	_ = ftl.AddLocaleFile("en_EU", path)

	val, err := ftl.LoadOneDynamic("", "en_EU", "missing_key")
	if err == nil {
		t.Fatalf("expected error for missing key, got value: %s", val)
	}
}

func TestSaveDynamicOverwritesExistingKeys(t *testing.T) {
	path, cleanup := createTempFile(t, "greet = hello\n")
	defer cleanup()

	ftl := source.NewFtl()
	_ = ftl.AddLocaleFile("en_EU", path)

	data := []source.Object{
		{LocaleCode: "en_EU", Key: "greet", Value: "hi"},
	}
	err := ftl.SaveDynamic("", data)
	if err != nil {
		t.Fatalf("SaveDynamic error: %v", err)
	}

	val, err := ftl.LoadOneDynamic("", "en_EU", "greet")
	if err != nil {
		t.Fatalf("LoadOneDynamic error: %v", err)
	}
	if val != "hi" {
		t.Errorf("expected 'hi', got '%s'", val)
	}
}

func TestFormatFTLEntrySerializesMultilineValues(t *testing.T) {
	value := "first line\nsecond line"

	got := source.FormatFTLEntry("message", value)
	want := "message =\n    first line\n    second line\n"

	if got != want {
		t.Fatalf("FormatFTLEntry() = %q, want %q", got, want)
	}
}

func TestSaveDynamicPreservesMultilineValues(t *testing.T) {
	path, cleanup := createTempFile(t, "")
	defer cleanup()

	ftl := source.NewFtl()
	_ = ftl.AddLocaleFile("en_EU", path)

	value := strings.Join([]string{
		"first line",
		"line with equals = still value",
		"# markdown heading, not an FTL comment",
		"last line",
	}, "\n")

	err := ftl.SaveDynamic("", []source.Object{
		{LocaleCode: "en_EU", Key: "complex", Value: value},
	})
	if err != nil {
		t.Fatalf("SaveDynamic error: %v", err)
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}

	expectedEntry := "complex =\n" +
		"    first line\n" +
		"    line with equals = still value\n" +
		"    # markdown heading, not an FTL comment\n" +
		"    last line\n"
	if !strings.Contains(string(raw), expectedEntry) {
		t.Fatalf("saved FTL does not contain formatted multiline entry:\n%s", string(raw))
	}

	got, err := ftl.LoadOneDynamic("", "en_EU", "complex")
	if err != nil {
		t.Fatalf("LoadOneDynamic error: %v", err)
	}
	if got != value {
		t.Fatalf("LoadOneDynamic() = %q, want %q", got, value)
	}
}
