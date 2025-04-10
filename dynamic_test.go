package word

import (
	"testing"
	"time"

	"github.com/summit-fi/wordsdk-go/source"
)

func remoteConnection() SDK {
	config := &Config{

		Source: source.NewRemote(
			"http://localhost:8000/api/v1",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkaWQiOiI0MmM5MTE1NC0wZTJiLTExZjAtODhkNS0yY2YwNWQ1N2EzZjIiLCJleHAiOjE3NzQ5NTk5NDEsImtleSI6ImEwMTM1NDQ3NzNiNDRkM2M4YmQ1YWMwYWZjYTY0MDAwIiwicGlkIjoiZTgyNTFmNzItYjE4MS0xMWVmLTkwNDYtNmE0YmQxZGQ3MzBjIiwidWlkIjoiOWQ1ZGEwMGEtYjFiNy0xMWVmLWE4YzctNmE0YmQxZGQ3MzBiIn0.0cboeSjqja6VZVNT4TnSqhG00Xwl_pfEeBH6F__Hz88"),
		UpdateInterval: 10 * time.Second,
		MaxCacheSizeMB: 256,
	}
	word, err := NewClient(config)
	if err != nil {
		panic(err)
	}
	word.SetLogger(&DefaultLogger{
		LogLevel: LogLevelDebug,
	})

	return word
}

func TestDynamicContent_T(t *testing.T) {
	connect := remoteConnection()

	// Test 1
	connect = connect.EnableDynamicContent(XKeyGen("S", "summit", "notification"))
	t.Logf("X-Dynamic-Key: %s", XKeyGen("S", "summit", "notification"))

	err := connect.Dynamic().SaveTranslation("en_EU", "core.every", "Everything is ok")
	if err != nil {
		t.Errorf("Failed to save translation: %v", err)
		return
	}

	connect = connect.EnableDynamicContent(XKeyGen("S", "summit", "notification"))
	str1 := connect.Dynamic().T("en_EU", "core.every")
	if str1 == "core.every" {
		t.Errorf("Failed to save translation: %v", str1)
		return
	}
	t.Logf("Translation: %s", str1)

}

func TestDynamicContent_TA(t *testing.T) {
	connect := remoteConnection()

	connect = connect.EnableDynamicContent(XKeyGen("S", "summit", "notification"))
	t.Logf("X-Dynamic-Key: %s", XKeyGen("S", "summit", "notification"))

	err := connect.Dynamic().SaveTranslation("en_EU", "select.ICU.test", `{every, select, weekly {Every week} monthly {Every month} yearly {Every year} other {Never}}`)
	if err != nil {
		t.Errorf("Failed to save translation: %v", err)
		return
	}

	str := connect.Dynamic().TA("en_EU", "select.ICU.test", map[string]interface{}{
		"every": "weekly",
	})

	if str == "plural.test" {
		t.Errorf("Failed to save translation: %v", str)
		return
	}

	t.Logf("Translation: %s", str)
}

func TestDynamicContent_SaveTranslations(t *testing.T) {
	connect := remoteConnection()

	connect = connect.EnableDynamicContent(XKeyGen("S", "summit", "notification"))
	t.Logf("X-Dynamic-Key: %s", XKeyGen("S", "summit", "notification"))

	err := connect.Dynamic().SaveTranslations([]source.Object{
		{
			LocaleCode: "en_EU",
			Key:        "core.every1",
			Value:      "Every week",
		},
		{
			LocaleCode: "en_EU",
			Key:        "ordinal.ICU.test",
			Value:      `Our {count, selectordinal, one {#st} two {#nd} few {#rd} other {#th}} tree!`,
		},
	})

	if err != nil {
		t.Errorf("Failed to save translations: %v", err)
		return
	}

	str := connect.Dynamic().TA("en_EU", "ordinal.ICU.test", map[string]interface{}{
		"count": 5,
	})
	if str == "select.ICU.test" {
		t.Errorf("Failed to save translation: %v", str)
		return
	}
	t.Logf("Translation: %s", str)
}
