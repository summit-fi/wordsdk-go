package word

import (
	"fmt"
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

	return word
}
func TestDynamicContent_T(t *testing.T) {
	connect := remoteConnection()

	// Test 1
	connect = connect.EnableDynamicContent(XKeyGen("S", "summit", "notification"))
	//connect.Dynamic().T("en_US", "core.every")
	fmt.Println(XKeyGen("S", "summit", "notification"))

	err := connect.Dynamic().SaveTranslation("en_EU", "core.every", "Everything is ok")
	if err != nil {
		t.Errorf("Failed to save translation: %v", err)
		return
	}
	//// Test 2
	connect = connect.EnableDynamicContent(XKeyGen("S", "summit", "notification"))
	str := connect.Dynamic().T("en_EU", "core.every")
	if str == "core.every" {
		t.Errorf("Failed to save translation: %v", str)
		return
	}

	fmt.Println(str, len(str))
}
