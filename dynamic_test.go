package word

import (
	"testing"

	"github.com/summit-fi/wordsdk-go/source"
)

func TestRemoteDynamicContent_T(t *testing.T) {
	connect, err := ftlClientWithSaveStrategy(SaveStrategyOnDemand)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	// Test 1
	connect = connect.EnableDynamicContent(XKeyGen("S", "summit", "notification"))
	t.Logf("X-Dynamic-Key: %s", XKeyGen("S", "summit", "notification"))

	err = connect.Dynamic().SaveTranslation("en_EU", "core_every", "Everything is ok")
	if err != nil {
		t.Errorf("Failed to save translation: %v", err)
		return
	}

	connect = connect.EnableDynamicContent(XKeyGen("S", "summit", "notification"))
	str1 := connect.Dynamic().T("en_EU", "core_every")
	if str1 == "core.every" {
		t.Errorf("Failed to save translation: %v", str1)
		return
	}
	t.Logf("Translation: %s", str1)

}

func TestRemoteDynamicContent_TA(t *testing.T) {
	connect, err := ftlClientWithSaveStrategy(SaveStrategyOnDemand)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	connect = connect.EnableDynamicContent(XKeyGen("S", "summit", "notification"))

	value := `Test`
	err = connect.Dynamic().SaveTranslation("en_EU", "key", value)
	if err != nil {
		t.Errorf("Failed to save translation: %v", err)
		return
	}

	str := connect.Dynamic().TA("en_EU", XKeyGen("S", "summit", "notification"), map[string]interface{}{
		"every": "weekly",
	})

	if str == "plural.test" {
		t.Errorf("Failed to save translation: %v", str)
		return
	}

	t.Logf("Translation: %s", str)
}

func TestRemoteDynamicContent_SaveTranslations(t *testing.T) {
	connect, err := ftlClientWithSaveStrategy(SaveStrategyOnDemand)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	connect = connect.EnableDynamicContent(XKeyGen("S", "summit", "schedule"))
	t.Logf("X-Dynamic-Key: %s", XKeyGen("S", "summit", "schedule"))

	err = connect.Dynamic().SaveTranslations([]source.Object{
		{
			LocaleCode: "en_EU",
			Key:        "test_test",
			Value:      "Testing dynamic content",
		},
	})

	if err != nil {
		t.Errorf("Failed to save translations: %v", err)
		return
	}

	str := connect.Dynamic().T("en_EU", "test_test")
	if str == "test_test" {
		t.Errorf("Failed to retrieve translation: %v", str)
		return
	}

	t.Logf("Translation: %s", str)
}

func TestRemoteDynamicContent_SaveTranslation(t *testing.T) {
	connect, err := ftlClientWithSaveStrategy(SaveStrategyOnDemand)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	connect = connect.EnableDynamicContent(XKeyGen("S", "summit", "schedule"))
	t.Logf("X-Dynamic-Key: %s", XKeyGen("S", "summit", "schedule"))

	err = connect.Dynamic().SaveTranslation("en_EU", "test_test_2", "Testing dynamic content 2")
	if err != nil {
		t.Errorf("Failed to save translation: %v", err)
		return
	}

	str := connect.Dynamic().TA("en_EU", "test_test_2", map[string]any{})
	if str == "test_test_2" {
		t.Errorf("Failed to retrieve translation: %v", str)
		return
	}

	t.Logf("Translation: %s", str)
}
