package main

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
	return current[len(gopath+"/src/"):] + "/" + name, nil
}
