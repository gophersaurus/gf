package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path"
)

func project(name, git string, verbose bool) error {

	var org string

	// Check if we a name or git url.
	if len(name) == 0 && len(git) == 0 {
		return errors.New("Need a project name or git repository URL.")
	}

	// Extract the name of the project from the git repository URL.
	if len(name) == 0 && len(git) > 0 {
		base := path.Base(git)
		ext := path.Ext(base)
		name = base[0 : len(base)-len(ext)]
	}

	// If the git URL was provided, compare the org name to the current directory.
	// If the two are different, display a warning for the sake of $GOPATH.
	//
	// Currently this only supports github.
	if len(git) > 0 {

		uri := git

		if uri[0:len("git@github.com:")] == "git@github.com:" {
			uri = uri[len("git@github.com:"):]
		}

		u, err := url.Parse(uri)
		if err != nil {
			return err
		}

		dir, _ := path.Split(u.Path)
		if dir[0] == '/' {
			dir = dir[1:]
		}
		dir = path.Clean(dir)

		// Get the current working directory.
		wd, err := os.Getwd()
		if err != nil {
			return err
		}

		if wd != dir {
			fmt.Println("Warning: The current directory might not match the git URL provided.")
			fmt.Println("Do you want to continue and use " + dir + " instead of " + wd + " for your import paths?")
			if !askForConfirmation() {
				os.Exit(1)
			}
		}

		org = dir
	}

	// Clone the Gophersaurus repository.
	err := exec.Command("git", "clone", "https://git.target.com/gophersaurus/gophersaurus.git").Run()
	if err != nil {
		return err
	}

	fmt.Println("cloned https://git.target.com/gophersaurus/gophersaurus.git")

	// Rename gophersaurus directory by project name.
	err = exec.Command("mv", "gophersaurus", name).Run()
	if err != nil {
		return err
	}

	fmt.Println("renamed gophersaurus -> " + name)

	// Find and replace gophersaurus/gophersaurus, with org/name in the entire directory.
	replacement := map[string]string{"git.target.com/gophersaurus/gophersaurus": "github.com/" + org + "/" + name}
	if err = rwImports(name, replacement); err != nil {
		return err
	}

	fmt.Println("renamed gophersaurus -> " + name)

	return nil
}
