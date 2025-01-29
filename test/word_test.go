package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/summit-fi/wordsdk-go"
)

func TestWord_Translate(t *testing.T) {

	w := wordsdk_go.New(wordsdk_go.Config{})

	tr, err := w.Translate("spot", wordsdk_go.Custom(map[string]interface{}{"center": "hotel"}), wordsdk_go.Count(1))
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "Welcome to your one room.", tr)

	tr, err = w.Translate("spot", wordsdk_go.Custom(map[string]interface{}{"center": "tennis"}), wordsdk_go.Count(2)).String()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "Welcome to your 2 courts.", tr)

	tr, err = w.Translate("spot", wordsdk_go.Count(1))
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "Welcome to your one room.", tr)

	tr, err = w.Translate("cldr-month-narrow", wordsdk_go.Index(0))
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "J", tr)

}

func TestWord_Time(t *testing.T) {

}
