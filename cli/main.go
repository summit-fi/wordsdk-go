package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/summit-fi/wordsdk-go/source"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "wordsdk-go",
		Short: "Word SDK command line tools",
	}

	rootCmd.AddCommand(exportCmd())

	// Enable automatic environment variables
	viper.AutomaticEnv()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func exportCmd() *cobra.Command {
	var outputDir string

	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export translations to FTL files",
		RunE: func(cmd *cobra.Command, args []string) error {
			apiBase := viper.GetString("api_base")
			accessKey := viper.GetString("access_key")

			if apiBase == "" || accessKey == "" {
				return fmt.Errorf("api_base and access_key must be provided via flags or env")
			}

			src := source.NewRemote(apiBase, accessKey)
			data, _, err := src.LoadAllStatic("")
			if err != nil {
				return err
			}

			absOut, err := filepath.Abs(outputDir)
			if err != nil {
				return err
			}
			if err := os.MkdirAll(absOut, 0o755); err != nil {
				return err
			}

			byLocale := map[string][]source.Object{}
			for _, obj := range data {
				byLocale[obj.LocaleCode] = append(byLocale[obj.LocaleCode], obj)
			}

			for locale, objs := range byLocale {
				filePath := filepath.Join(absOut, fmt.Sprintf("%s.ftl", locale))
				f, err := os.Create(filePath)
				if err != nil {
					return err
				}
				var b strings.Builder
				for _, o := range objs {
					b.WriteString(o.Key)
					b.WriteString(" = ")
					b.WriteString(o.Value)
					if !strings.HasSuffix(o.Value, "\n") {
						b.WriteString("\n")
					}
				}
				if _, err := f.WriteString(b.String()); err != nil {
					f.Close()
					return err
				}
				f.Close()
				fmt.Fprintln(cmd.OutOrStdout(), "exported", filePath)
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&outputDir, "output", "o", ".", "directory to save exported files")
	viper.BindPFlag("output", cmd.Flags().Lookup("output"))

	cmd.Flags().String("api-base", "", "API base URL")
	viper.BindPFlag("api_base", cmd.Flags().Lookup("api-base"))
	cmd.Flags().String("access-key", "", "API access key")
	viper.BindPFlag("access_key", cmd.Flags().Lookup("access-key"))

	return cmd
}
