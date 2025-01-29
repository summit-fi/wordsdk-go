package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/summit-fi/wordsdk-go"
	"golang.org/x/text/language"
)

func TestWord_Translate(t *testing.T) {

	w, err := wordsdk_go.New(wordsdk_go.Config{
		FTLFilePath: []string{"./test.ftl"},
	})
	if err != nil {
		t.Error(err)
	}

	tr, err := w.Translate(
		language.English,
		"spot",
		wordsdk_go.Custom(map[string]interface{}{"center": "hotel"}),
		wordsdk_go.Count(1))

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Welcome to your one room.", tr)

	tr, err = w.Translate(
		language.English,
		"spot",
		wordsdk_go.Custom(map[string]interface{}{"center": "tennis"}),
		wordsdk_go.Count(2))

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Welcome to your 2 courts.", tr)

	tr, err = w.Translate(
		language.English,
		"spot",
		wordsdk_go.Count(1))

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Welcome to your one room.", tr)

	tr, err = w.Translate(
		language.English,
		"cldr-month-narrow",
		wordsdk_go.Index(0))

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "J", tr)

}

func TestWord_Time(t *testing.T) {
	w, err := wordsdk_go.New(wordsdk_go.Config{
		FTLFilePath: []string{"./test.ftl"},
	})
	if err != nil {
		t.Error(err)
	}

	tr, err := w.Translate(language.English, "function-test", wordsdk_go.Argument("date", "2025-01-01"))
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "2025 01 01", tr)
}
