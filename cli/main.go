package main

import (
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
	rootCmd.PersistentFlags().String("api-key", "", "API key")
	rootCmd.PersistentFlags().String("dynamic-key", "", "Dynamic key for fetching dynamic translations")
	rootCmd.PersistentFlags().String("environment", "production", "API environment: production or stage")
	rootCmd.PersistentFlags().StringP("output", "o", "./exported", "Output directory")
}
