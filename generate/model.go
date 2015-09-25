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

func Model(args []string) error {

	if len(args) < 1 {
		return errors.New("expected to receive the model name as an argument")
	}

	// get the model name
	name := args[0]

	// create the OS agnostic model path
	file := filepath.Join("app", "models", name+".go")

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
	t, err := template.New("model").Parse(templates.Model)
	if err != nil {
		return err
	}

	// write it to disk
	t.Execute(f, name)
	return nil
}
