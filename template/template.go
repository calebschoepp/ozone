package template

import (
	"embed"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/calebschoepp/ozone/directive"
)

//go:embed primary
var primaryTmpls embed.FS

var outDir = ".ozone/"

// Template for the primary CLI type
// TODO give better name than Primary
type primary struct {
	D    directive.Directive
	wasm []string
}

// Construct a new primary CLI template
func NewPrimary(d directive.Directive, wasm []string) primary {
	return primary{
		D:    d,
		wasm: wasm,
	}
}

// Templates out the source for a new primary CLI into the given path
func (p *primary) Template(path string) error {
	err := fs.WalkDir(primaryTmpls, ".", p.handleWalk)
	if err != nil {
		return err
	}

	for _, wasm := range p.wasm {
		err = copy(wasm, outDir+"cmd/"+filepath.Base(wasm))
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *primary) handleWalk(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if d.IsDir() {
		return nil
	}

	templateFile, err := primaryTmpls.ReadFile(path)
	if err != nil {
		return err
	}

	template, err := template.New(d.Name()).Parse(string(templateFile))
	if err != nil {
		return err
	}

	outFilePath := outDir + strings.TrimSuffix(strings.TrimPrefix(path, "primary"), ".tmpl")
	err = os.MkdirAll(filepath.Dir(outFilePath), 0775)
	if err != nil {
		return err
	}

	outFile, err := os.Create(outFilePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	if d.Name() == "cmd.go.tmpl" {
		for _, cmd := range p.D.Commands {
			err = template.Execute(outFile, cmd)
			if err != nil {
				return err
			}
		}
	} else {
		err = template.Execute(outFile, *p)
		if err != nil {
			return err
		}
	}

	return nil
}

func copy(src, dest string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	return nil
}
