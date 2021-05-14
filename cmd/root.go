package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "ozone",
	Short:   "Ozone is a declarative and language-agnostic CLI framework",
	Long:    "Ozone is a CLI framework that uses the power of Web Assembly to offer a declarative and language-agnostic CLI framework",
	Version: "v0.0.1",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(cleanCmd)
}
