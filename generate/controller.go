package generate

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"text/template"
	"unicode"
	"unicode/utf8"

	"github.com/gophersaurus/gf/generate/templates"
)

func Controller(args []string) error {

	if len(args) < 1 {
		return errors.New("expected to receive the controller name as an argument")
	}

	// get the controller name
	name := args[0]

	// create the OS agnostic controller path
	file := filepath.Join("app", "controllers", name+".go")

	// uppercase the first rune in name string
	r, size := utf8.DecodeRuneInString(name)
	R := unicode.ToUpper(r)
	name = string(R) + name[size:]

	// ensure directories exist
	if err := os.MkdirAll(path.Dir(file), 0777); err != nil {
		return err
	}

	// create the file
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	// parse template
	t, err := template.New("controller").Parse(templates.Controller)
	if err != nil {
		return err
	}

	// write it to disk
	t.Execute(f, name)
	return nil
}
