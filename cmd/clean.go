package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove intermediate build files",
	Long:  "Remove all intermediate build files that Ozone uses to generate your CLI binary.",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
		fmt.Println("Cleaning")
	},
}
