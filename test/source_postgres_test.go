package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/summit-fi/wordsdk-go/source"
)

const (
	postgresURL = "postgresql://postgres:*****@localhost:5432/your_database_name?sslmode=disable"
)

func TestSourcePostgres_SaveDynamic(t *testing.T) {
	ctx := context.TODO()
	src, err := source.NewPostgres(ctx, postgresURL, nil)
	assert.Nil(t, err, "NewPostgres should not return an error")

	// Test SaveDynamic
	data := []source.Object{
		{LocaleCode: "en_EU", Key: "greet", Value: "hello"},
		{LocaleCode: "en_EU", Key: "farewell", Value: "goodbye"},
	}
	err = src.SaveDynamic("", data)
	assert.Nil(t, err, "SaveDynamic should not return an error")

	// Verify that the data was saved correctly
	translations, err := src.LoadOneDynamic("", "en_EU", "greet")
	assert.Nil(t, err, "LoadAllStatic should not return an error")
	assert.Contains(t, translations, "hello", "Translations should contain the saved 'greet' key")
	translations, err = src.LoadOneDynamic("", "en_EU", "farewell")
	assert.Nil(t, err, "LoadAllStatic should not return an error")
	assert.Contains(t, translations, "goodbye", "Translations should contain the saved 'farewell' key")
}
