package test

import (
	"os"
	"strings"
	"testing"

	"github.com/summit-fi/wordsdk-go/fluent"
	"github.com/summit-fi/wordsdk-go/fluent/cldr"
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

func TestFtlParseWithErrorParsesFluentEntriesInSourceOrder(t *testing.T) {
	data := []byte(`# A regular Fluent comment
welcome = Hello, { $name }
    Welcome back.
    .title = Greeting

-brand = Summit

count =
    { $count ->
        [one] One item
       *[other] { $count } items
    }
`)

	objects, err := source.FtlParseWithError("en_EU", data)
	if err != nil {
		t.Fatalf("FtlParseWithError error: %v", err)
	}
	if len(objects) != 3 {
		t.Fatalf("FtlParseWithError returned %d objects, want 3: %#v", len(objects), objects)
	}

	wantKeys := []string{"welcome", "-brand", "count"}
	for index, want := range wantKeys {
		if objects[index].Key != want {
			t.Fatalf("object %d key = %q, want %q", index, objects[index].Key, want)
		}
	}
	if !strings.Contains(objects[0].Value, "{ $name }") {
		t.Fatalf("welcome value lost its placeable: %q", objects[0].Value)
	}
	if !strings.Contains(objects[0].Value, ".title = Greeting") {
		t.Fatalf("welcome value lost its attribute: %q", objects[0].Value)
	}
	if !strings.Contains(objects[2].Value, "$count ->") {
		t.Fatalf("count value lost its selector: %q", objects[2].Value)
	}

	for _, object := range objects {
		entry := source.FormatFTLEntry(object.Key, object.Value)
		if _, err := source.FtlParseWithError("en_EU", []byte(entry)); err != nil {
			t.Fatalf("serialized object %q is invalid FTL: %v\n%s", object.Key, err, entry)
		}
	}
}

func TestFtlParseWithErrorReportsMalformedFTL(t *testing.T) {
	data := []byte("valid = first\ninvalid key = second\nafter = third\n")

	objects, err := source.FtlParseWithError("en_EU", data)
	if err == nil {
		t.Fatal("FtlParseWithError returned no error for malformed FTL")
	}
	if !strings.Contains(err.Error(), "line 2") {
		t.Fatalf("parse error does not identify the malformed line: %v", err)
	}
	if len(objects) != 2 || objects[0].Key != "valid" || objects[1].Key != "after" {
		t.Fatalf("valid entries were not recovered in source order: %#v", objects)
	}
}

func TestSaveDynamicUpdatesEntryWithoutDuplicatingIt(t *testing.T) {
	path, cleanup := createTempFile(t, "# keep this comment\nmessage = old\n    .title = Keep this attribute\nother = untouched\n")
	defer cleanup()

	ftl := source.NewFtl()
	if err := ftl.AddLocaleFile("en_EU", path); err != nil {
		t.Fatalf("AddLocaleFile error: %v", err)
	}

	for _, value := range []string{"new", "newer"} {
		if err := ftl.SaveDynamic("", []source.Object{{
			LocaleCode: "en_EU",
			Key:        "message",
			Value:      value,
		}}); err != nil {
			t.Fatalf("SaveDynamic(%q) error: %v", value, err)
		}
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}
	content := string(raw)
	if strings.Count(content, "message =") != 1 {
		t.Fatalf("message was duplicated:\n%s", content)
	}
	for _, preserved := range []string{"# keep this comment", ".title = Keep this attribute", "other = untouched"} {
		if !strings.Contains(content, preserved) {
			t.Fatalf("SaveDynamic removed %q:\n%s", preserved, content)
		}
	}

	got, err := ftl.LoadOneDynamic("", "en_EU", "message")
	if err != nil {
		t.Fatalf("LoadOneDynamic error: %v", err)
	}
	if got != "newer" {
		t.Fatalf("LoadOneDynamic() = %q, want %q", got, "newer")
	}
	if _, err := source.FtlParseWithError("en_EU", raw); err != nil {
		t.Fatalf("SaveDynamic produced invalid FTL: %v\n%s", err, content)
	}
}

func TestSaveDynamicRejectsInvalidUpdateWithoutChangingFile(t *testing.T) {
	const initial = "message = unchanged\n"
	path, cleanup := createTempFile(t, initial)
	defer cleanup()

	ftl := source.NewFtl()
	if err := ftl.AddLocaleFile("en_EU", path); err != nil {
		t.Fatalf("AddLocaleFile error: %v", err)
	}

	err := ftl.SaveDynamic("", []source.Object{{
		LocaleCode: "en_EU",
		Key:        "invalid key",
		Value:      "new value",
	}})
	if err == nil {
		t.Fatal("SaveDynamic returned no error for an invalid key")
	}

	raw, readErr := os.ReadFile(path)
	if readErr != nil {
		t.Fatalf("ReadFile error: %v", readErr)
	}
	if string(raw) != initial {
		t.Fatalf("file changed after rejected update: %q", raw)
	}
}

func TestSaveDynamicEmptyValueRoundTrips(t *testing.T) {
	path, cleanup := createTempFile(t, "message = old\n")
	defer cleanup()

	ftl := source.NewFtl()
	if err := ftl.AddLocaleFile("en_EU", path); err != nil {
		t.Fatalf("AddLocaleFile error: %v", err)
	}
	if err := ftl.SaveDynamic("", []source.Object{{
		LocaleCode: "en_EU",
		Key:        "message",
		Value:      "",
	}}); err != nil {
		t.Fatalf("SaveDynamic error: %v", err)
	}

	got, err := ftl.LoadOneDynamic("", "en_EU", "message")
	if err != nil {
		t.Fatalf("LoadOneDynamic error: %v", err)
	}
	if got != "" {
		t.Fatalf("LoadOneDynamic() = %q, want an empty string", got)
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}
	if _, err := source.FtlParseWithError("en_EU", raw); err != nil {
		t.Fatalf("empty value was not encoded as valid FTL: %v", err)
	}
}

func TestSaveDynamicAcceptsParsedEntryWithAttributes(t *testing.T) {
	const initial = "button = Old\n    .title = Old title\n"
	path, cleanup := createTempFile(t, initial)
	defer cleanup()

	parsed, err := source.FtlParseWithError("en_EU", []byte("button = New\n    .title = New title\n"))
	if err != nil {
		t.Fatalf("FtlParseWithError error: %v", err)
	}
	ftl := source.NewFtl()
	if err := ftl.AddLocaleFile("en_EU", path); err != nil {
		t.Fatalf("AddLocaleFile error: %v", err)
	}
	if err := ftl.SaveDynamic("", parsed); err != nil {
		t.Fatalf("SaveDynamic rejected parsed FTL: %v", err)
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}
	content := string(raw)
	for _, want := range []string{"button =", "New", ".title = New title"} {
		if !strings.Contains(content, want) {
			t.Fatalf("updated FTL does not contain %q:\n%s", want, content)
		}
	}
	if strings.Contains(content, "Old") {
		t.Fatalf("old value or attributes remain after full entry update:\n%s", content)
	}
	if _, err := source.FtlParseWithError("en_EU", raw); err != nil {
		t.Fatalf("SaveDynamic produced invalid FTL: %v\n%s", err, content)
	}
}

func TestSaveDynamicAcceptsParsedAttributeOnlyEntry(t *testing.T) {
	const entry = "button =\n    .title = New title\n"
	path, cleanup := createTempFile(t, "button = Old\n")
	defer cleanup()

	parsed, err := source.FtlParseWithError("en_EU", []byte(entry))
	if err != nil {
		t.Fatalf("FtlParseWithError error: %v", err)
	}
	ftl := source.NewFtl()
	if err := ftl.AddLocaleFile("en_EU", path); err != nil {
		t.Fatalf("AddLocaleFile error: %v", err)
	}
	if err := ftl.SaveDynamic("", parsed); err != nil {
		t.Fatalf("SaveDynamic rejected attribute-only FTL: %v", err)
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}
	objects, err := source.FtlParseWithError("en_EU", raw)
	if err != nil {
		t.Fatalf("SaveDynamic produced invalid FTL: %v\n%s", err, raw)
	}
	if len(objects) != 1 || !strings.Contains(objects[0].Value, ".title = New title") {
		t.Fatalf("attribute-only entry did not round-trip: %#v\n%s", objects, raw)
	}
}

func TestSaveDynamicAcceptsParsedTermWithAttributes(t *testing.T) {
	const initial = "-brand = Old\n    .title = Old title\n"
	path, cleanup := createTempFile(t, initial)
	defer cleanup()

	parsed, err := source.FtlParseWithError("en_EU", []byte("-brand = New\n    .title = New title\n"))
	if err != nil {
		t.Fatalf("FtlParseWithError error: %v", err)
	}
	ftl := source.NewFtl()
	if err := ftl.AddLocaleFile("en_EU", path); err != nil {
		t.Fatalf("AddLocaleFile error: %v", err)
	}
	if err := ftl.SaveDynamic("", parsed); err != nil {
		t.Fatalf("SaveDynamic rejected term with attributes: %v", err)
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}
	content := string(raw)
	if strings.Contains(content, "--brand") || strings.Contains(content, "Old") {
		t.Fatalf("term was replaced incorrectly:\n%s", content)
	}
	if !strings.Contains(content, "-brand =") || !strings.Contains(content, ".title = New title") {
		t.Fatalf("term or attribute is missing:\n%s", content)
	}
	if _, err := source.FtlParseWithError("en_EU", raw); err != nil {
		t.Fatalf("SaveDynamic produced invalid FTL: %v\n%s", err, content)
	}
}

func TestSaveDynamicPreservesCRLFAndSymlink(t *testing.T) {
	directory := t.TempDir()
	target := directory + "/target.ftl"
	link := directory + "/locale.ftl"
	if err := os.WriteFile(target, []byte("message = old\r\n"), 0640); err != nil {
		t.Fatalf("WriteFile error: %v", err)
	}
	if err := os.Symlink(target, link); err != nil {
		t.Fatalf("Symlink error: %v", err)
	}

	ftl := source.NewFtl()
	if err := ftl.AddLocaleFile("en_EU", link); err != nil {
		t.Fatalf("AddLocaleFile error: %v", err)
	}
	if err := ftl.SaveDynamic("", []source.Object{{
		LocaleCode: "en_EU",
		Key:        "message",
		Value:      "new\nvalue",
	}}); err != nil {
		t.Fatalf("SaveDynamic error: %v", err)
	}

	info, err := os.Lstat(link)
	if err != nil {
		t.Fatalf("Lstat error: %v", err)
	}
	if info.Mode()&os.ModeSymlink == 0 {
		t.Fatal("SaveDynamic replaced the locale symlink")
	}
	raw, err := os.ReadFile(target)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}
	withoutCRLF := strings.ReplaceAll(string(raw), "\r\n", "")
	if strings.ContainsAny(withoutCRLF, "\r\n") {
		t.Fatalf("SaveDynamic introduced mixed line endings: %q", raw)
	}
}

func TestFtlParseWithErrorRejectsInvalidUTF8AndDuplicateKeys(t *testing.T) {
	for name, data := range map[string][]byte{
		"invalid UTF-8": {0xff, '=', 'x'},
		"duplicate key": []byte("message = first\nmessage = second\n"),
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := source.FtlParseWithError("en_EU", data); err == nil {
				t.Fatalf("FtlParseWithError accepted %s", name)
			}
		})
	}
}

func TestSaveDynamicRejectsInvalidUTF8WithoutChangingFile(t *testing.T) {
	for name, testCase := range map[string]struct {
		initial []byte
		update  source.Object
	}{
		"source": {
			initial: []byte{'m', 'e', 's', 's', 'a', 'g', 'e', ' ', '=', ' ', 0xff, '\n'},
			update:  source.Object{LocaleCode: "en_EU", Key: "message", Value: "new"},
		},
		"update": {
			initial: []byte("message = old\n"),
			update:  source.Object{LocaleCode: "en_EU", Key: "message", Value: string([]byte{0xff})},
		},
	} {
		t.Run(name, func(t *testing.T) {
			path, cleanup := createTempFile(t, string(testCase.initial))
			defer cleanup()

			ftl := source.NewFtl()
			if err := ftl.AddLocaleFile("en_EU", path); err != nil {
				t.Fatalf("AddLocaleFile error: %v", err)
			}
			if err := ftl.SaveDynamic("", []source.Object{testCase.update}); err == nil {
				t.Fatal("SaveDynamic accepted invalid UTF-8")
			}
			raw, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("ReadFile error: %v", err)
			}
			if string(raw) != string(testCase.initial) {
				t.Fatalf("file changed after rejected update: %q", raw)
			}
		})
	}
}

func TestAddLocaleFileRejectsSameFileForDifferentLocales(t *testing.T) {
	path, cleanup := createTempFile(t, "message = value\n")
	defer cleanup()

	ftl := source.NewFtl()
	if err := ftl.AddLocaleFile("en_EU", path); err != nil {
		t.Fatalf("AddLocaleFile(en_EU) error: %v", err)
	}
	if err := ftl.AddLocaleFile("fr_FR", path); err == nil {
		t.Fatal("AddLocaleFile accepted the same file for two locales")
	}
}

func TestFtlParseWithErrorRoundTripsThroughBundle(t *testing.T) {
	data := []byte(`welcome = Hello, { $name }
    Welcome back.
    .title = Greeting

count =
    { $count ->
        [one] One item
       *[other] { $count } items
    }
`)
	objects, err := source.FtlParseWithError("en_EU", data)
	if err != nil {
		t.Fatalf("FtlParseWithError error: %v", err)
	}

	bundle := fluent.NewBundle(cldr.LanguageEnUS)
	for _, object := range objects {
		// Match Client.UpdateBundle's construction of a resource from Object.
		resource, parseErrors := fluent.NewResource(object.Key + " = " + object.Value + "\n")
		if len(parseErrors) > 0 {
			t.Fatalf("serialized object %q is invalid: %v", object.Key, parseErrors)
		}
		if addErrors := bundle.AddResource(resource); len(addErrors) > 0 {
			t.Fatalf("could not add object %q to bundle: %v", object.Key, addErrors)
		}
	}

	welcome, resolveErrors, err := bundle.FormatFullMessage(
		"welcome",
		fluent.WithVariable("name", "Ada"),
	)
	if err != nil || len(resolveErrors) > 0 {
		t.Fatalf("FormatFullMessage errors: %v, %v", err, resolveErrors)
	}
	if got := welcome.Text(""); got != "Hello, Ada\nWelcome back." {
		t.Fatalf("welcome text = %q", got)
	}
	if got := welcome.Attr("title", ""); got != "Greeting" {
		t.Fatalf("welcome title = %q", got)
	}

	count, resolveErrors, err := bundle.FormatMessage("count", fluent.WithVariable("count", 2))
	if err != nil || len(resolveErrors) > 0 {
		t.Fatalf("FormatMessage errors: %v, %v", err, resolveErrors)
	}
	if count != "2 items" {
		t.Fatalf("count message = %q, want %q", count, "2 items")
	}
}

func TestFtlParseWithErrorHandlesLargeLines(t *testing.T) {
	want := strings.Repeat("x", 128*1024)
	objects, err := source.FtlParseWithError("en_EU", []byte("large = "+want+"\n"))
	if err != nil {
		t.Fatalf("FtlParseWithError error: %v", err)
	}
	if len(objects) != 1 || objects[0].Value != want {
		t.Fatalf("large value was truncated: got %d bytes, want %d", len(objects[0].Value), len(want))
	}
}

func TestSaveDynamicValidatesAllLocalesBeforeWriting(t *testing.T) {
	firstPath, firstCleanup := createTempFile(t, "message = first\n")
	defer firstCleanup()
	secondPath, secondCleanup := createTempFile(t, "invalid key = second\n")
	defer secondCleanup()

	ftl := source.NewFtl()
	if err := ftl.AddLocaleFile("en_EU", firstPath); err != nil {
		t.Fatalf("AddLocaleFile(en_EU) error: %v", err)
	}
	if err := ftl.AddLocaleFile("fr_FR", secondPath); err != nil {
		t.Fatalf("AddLocaleFile(fr_FR) error: %v", err)
	}
	err := ftl.SaveDynamic("", []source.Object{
		{LocaleCode: "en_EU", Key: "message", Value: "updated"},
		{LocaleCode: "fr_FR", Key: "message", Value: "updated"},
	})
	if err == nil {
		t.Fatal("SaveDynamic returned no error for malformed second locale")
	}

	raw, readErr := os.ReadFile(firstPath)
	if readErr != nil {
		t.Fatalf("ReadFile error: %v", readErr)
	}
	if string(raw) != "message = first\n" {
		t.Fatalf("first locale changed before all locales were validated: %q", raw)
	}
}
