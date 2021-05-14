package directive

import (
	"errors"
	"fmt"

	"golang.org/x/mod/semver"
	"gopkg.in/yaml.v2"
)

// TODO Handle both single/multi command CLIs
// TODO Handle notion of Executables and CallableFns
// TODO Handle more robust notion of args i.e. more than just names as string to get a count

// Directive completely describes how a CLI should be configured and which wasm it should run
type Directive struct {
	Name       string    `yaml:"name"`
	AppVersion string    `yaml:"appVersion"`
	Commands   []Command `yaml:"commands"`
}

// Command represents the configuration of a single arbitrarily nested sub-command for the CLI
type Command struct {
	Name  string   `yaml:"name"`
	Args  []string `yaml:"args"`
	Steps []string `yaml:"steps"`
}

// Marshal outputs the YAML bytes of the Directive
func (d *Directive) Marshal() ([]byte, error) {
	return yaml.Marshal(d)
}

// Unmarshal unmarshals YAML bytes into a Directive struct
// it also calculates a map of FQFNs for later use
func (d *Directive) Unmarshal(in []byte) error {
	return yaml.Unmarshal(in, d)
}

// Validate validates a directive
func (d *Directive) Validate() error {
	problems := &problems{}

	if d.Name == "" {
		problems.add(errors.New("name is missing"))
	}

	if !semver.IsValid(d.AppVersion) {
		problems.add(errors.New("app version is not a valid semantic version"))
	}

	if len(d.Commands) < 1 {
		problems.add(errors.New("no commands listed"))
	}

	cmds := map[string]bool{}

	for i, c := range d.Commands {
		if c.Name == "" {
			problems.add(fmt.Errorf("command at position %d missing name", i))
			continue
		}

		if _, exists := cmds[c.Name]; exists {
			problems.add(fmt.Errorf("duplicate command %s found", c.Name))
			continue
		}
		cmds[c.Name] = true

		// TODO validate arguments

		if len(c.Steps) == 0 {
			problems.add(fmt.Errorf("command %s has no steps", c.Name))
			continue
		}

	}

	return problems.render()
}

// Convenience type to display all validation errors at once
type problems []error

func (p *problems) add(err error) {
	*p = append(*p, err)
}

func (p *problems) render() error {
	if len(*p) == 0 {
		return nil
	}

	text := fmt.Sprintf("found %d problems:", len(*p))

	for _, err := range *p {
		text += fmt.Sprintf("\n\t%s", err.Error())
	}

	return errors.New(text)
}
