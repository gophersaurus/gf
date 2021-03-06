package project

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/jbenet/go-os-rename"
)

// create creates a new Gopheraurus project.
func Create(name, git string, verbose, skip, skipgit bool) error {

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
	importpath, err := compare(name, git, skip)
	if err != nil {
		return err
	}
	cleanimportpath := filepath.ToSlash(importpath)

	if err := os.RemoveAll(".gf_tmp"); err != nil {
		return err
	}

	// Clone the Gophersaurus repository.
	if err := exec.Command("git", "clone", "https://github.com/gophersaurus/framework.git", ".gf_tmp").Run(); err != nil {
		fmt.Println()
		return err
	}

	// Rename gophersaurus directory by project name.
	if err := osrename.Rename(".gf_tmp", name); err != nil {
		return err
	}

	// Find and replace gophersaurus/gophersaurus, with org/name in the entire directory.
	replacement := map[string]string{"github.com/gophersaurus/framework": cleanimportpath}
	if err = rewrite(name, replacement); err != nil {
		return err
	}

	if !skipgit {

		// Set the git remote upstream for easy updating.
		if err := exec.Command("git", "-C", name, "remote", "add", "upstream", "https://github.com/gophersaurus/framework.git").Run(); err != nil {
			return err
		}

		// Check if a git URL has been provided.
		if len(git) > 0 {

			// set the git remote origin based on git URL provided.
			err = exec.Command("git", "-C", name, "remote", "set-url", "origin", git).Run()
			if err != nil {
				return err
			}

			return nil
		}

		// set the git remote origin based on go src file path.
		if err := exec.Command("git", "-C", name, "remote", "set-url", "origin", "https://"+cleanimportpath+".git").Run(); err != nil {
			return err
		}
	}

	return nil
}
