package project

import (
	"errors"
	"os"
	"path/filepath"
)

func src(name string) (string, error) {

	// Identify the current working directory.
	current, err := filepath.Abs(filepath.Dir(os.Args[0]))

	// Check for errors.
	if err != nil {
		return "", err
	}

	// Identify the local $GOPATH.
	gopath := os.Getenv("GOPATH")

	// Check the $GOPATH value is valid.
	if len(gopath) == 0 {
		return "", errors.New("gf error: $GOPATH not set")
	}

	// Use the $GOPATH to strip everything out, but the base git URL.
	gopathsrc := filepath.Join(gopath, "/src") + string(os.PathSeparator)
	return filepath.Join(current[len(gopathsrc):], name), nil
}
