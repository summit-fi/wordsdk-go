package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd.AddCommand(exportCmd)
	cobra.CheckErr(rootCmd.Execute())
}

var rootCmd = &cobra.Command{
	Use:   "wordsdk",
	Short: "Word SDK command line interface",
}

func init() {
	rootCmd.PersistentFlags().String("api-key", getenv("WORDSDK_API_KEY", ""), "API key")
	rootCmd.PersistentFlags().String("dynamic-key", getenv("WORDSDK_DYNAMIC_KEY", ""), "Dynamic key for fetching dynamic translations")
	rootCmd.PersistentFlags().String("environment", getenv("WORDSDK_ENVIRONMENT", "production"), "API environment: production or stage")
	rootCmd.PersistentFlags().StringP("output", "o", getenv("WORDSDK_OUTPUT", "./exported"), "Output directory")
}

func getenv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}
