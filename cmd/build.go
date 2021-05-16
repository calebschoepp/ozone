package cmd

import (
	"fmt"
	"os"

	"github.com/calebschoepp/ozone/directive"
	"github.com/calebschoepp/ozone/template"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build [dir]",
	Short: "Build your CLI",
	Long:  "Build takes in a Directive.yaml file and the associated Wasm and produces a CLI binary.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO Validate that the dir is actually a directory
		dir := "."
		if len(args) > 0 {
			dir = args[0]
		}

		// Read and validate directive
		bytes, err := os.ReadFile(dir + "/Directive.yaml")
		if err != nil {
			return fmt.Errorf("could not find Directive.yaml in %s", dir)
		}
		directive := directive.Directive{}
		err = directive.Unmarshal(bytes)
		if err != nil {
			return err
		}
		err = directive.Validate()
		if err != nil {
			return err
		}

		// Make sure we have all of the wasm
		var wasm []string
		for _, command := range directive.Commands {
			path := dir + command.Module + ".wasm"
			if _, err = os.Stat(path); err != nil {
				return fmt.Errorf("could not find module at %s", path)
			}
			wasm = append(wasm, path)

		}

		// Templates out the source code for the CLI
		primary := template.NewPrimary(directive, wasm)
		err = primary.Template("./.ozone/") // TODO improve this to a more sophisticated format
		if err != nil {
			return err
		}

		return nil
	},
}
