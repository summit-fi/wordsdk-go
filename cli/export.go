package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/summit-fi/wordsdk-go/source"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export translations to FTL files",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey, _ := cmd.Flags().GetString("api-key")
		dynamicKey, _ := cmd.Flags().GetString("dynamic-key")
		env, _ := cmd.Flags().GetString("environment")
		outDir, _ := cmd.Flags().GetString("output")

		if apiKey == "" {
			return fmt.Errorf("api-key is required")
		}

		var apiURL string
		switch env {
		case "production":
			apiURL = "https://wordapi.thesumm.it/api/v1"
		case "stage":
			apiURL = "https://dev.wordapi.thesumm.it/api/v1"
		default:
			apiURL = env
		}

		src := source.NewRemote(apiURL, apiKey)

		var data []source.Object
		var err error
		var path string
		if len(dynamicKey) == 0 {
			path = "static"
			data, _, err = src.LoadAllStatic("")
			if err != nil {
				return fmt.Errorf("failed to load static translations: %w", err)
			}
		} else {
			path = filepath.Join("dynamic", dynamicKey)
			data, _, err = src.LoadAllDynamic(dynamicKey, "")
			if err != nil {
				return fmt.Errorf("failed to load dynamic translations: %w", err)
			}
		}

		err = exportObjectsToFTLFiles(data, outDir, path)
		if err != nil {
			return fmt.Errorf("failed to export translations: %w", err)
		}

		return nil
	},
}

func exportObjectsToFTLFiles(objects []source.Object, outDir string, path string) error {
	groups := make(map[string][]source.Object)
	for _, obj := range objects {
		if len(obj.LocaleCode) > 0 {
			groups[obj.LocaleCode] = append(groups[obj.LocaleCode], obj)
		}
	}

	if outDir == "" {
		outDir = "."
	}
	if err := os.MkdirAll(filepath.Join(outDir, path), 0o755); err != nil {
		return err
	}

	for locale, objs := range groups {
		var b strings.Builder
		for _, o := range objs {
			b.WriteString(formatFTLEntry(o.Key, o.Value))
		}
		path := filepath.Join(outDir, path, locale+".ftl")
		if err := os.WriteFile(path, []byte(b.String()), 0o644); err != nil {
			return err
		}
		fmt.Printf("Exported %s\n", path)
	}
	return nil
}

const ftlBlockIndent = "    "

func formatFTLEntry(key, value string) string {
	value = normalizeFTLNewlines(value)
	if !strings.Contains(value, "\n") {
		return fmt.Sprintf("%s = %s\n", key, value)
	}

	var builder strings.Builder
	builder.WriteString(key)
	builder.WriteString(" =\n")
	for _, line := range strings.Split(value, "\n") {
		builder.WriteString(ftlBlockIndent)
		builder.WriteString(line)
		builder.WriteString("\n")
	}
	return builder.String()
}

func normalizeFTLNewlines(value string) string {
	value = strings.ReplaceAll(value, "\r\n", "\n")
	return strings.ReplaceAll(value, "\r", "\n")
}
