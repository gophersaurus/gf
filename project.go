package main

import (
	"errors"
	"os/exec"
	"path"
)

func project(name, git string, verbose bool) error {

	// Check if we a name or git url.
	if len(name) == 0 && len(git) == 0 {
		return errors.New("Need a project name or git repository URL.")
	}

	// If the name isn't given, then extract the name of the project from the
	// git repository URL.
	if len(name) == 0 && len(git) > 0 {
		base := path.Base(git)
		ext := path.Ext(base)
		name = base[0 : len(base)-len(ext)]
	}

	// compare the current directory path and git path.
	path, err := compare(name, git)
	if err != nil {
		return err
	}

	// Clone the Gophersaurus repository.
	err = exec.Command("git", "clone", "https://git.target.com/gophersaurus/gophersaurus.git").Run()
	if err != nil {
		return err
	}

	// Rename gophersaurus directory by project name.
	err = exec.Command("mv", "gophersaurus", name).Run()
	if err != nil {
		return err
	}

	// Find and replace gophersaurus/gophersaurus, with org/name in the entire directory.
	replacement := map[string]string{"git.target.com/gophersaurus/gophersaurus": path}
	if err = rewrite(name, replacement); err != nil {
		return err
	}

	return nil
}
