package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

// compare compares the git URL path to the go src directory path.
// If the two are different then warn the user before preceding.
func compare(name, git string) (string, error) {

	// get the current directory path.
	dir, err := src(name)
	if err != nil {
		return "", err
	}

	// If the git URL is not provided skip this step.
	if len(git) > 0 {

		// make a copy of the provided git url to work with.
		uri := git

		// check for ssh github url.
		// we are cheating, only supporting github + github enterprise for now.
		if uri[0:len("git@")] == "git@" {
			uri = uri[len("git@"):]
			ss := strings.Split(uri, ":")
			uri = ss[0] + "/" + ss[1]
		}

		// check for http or https github url.
		if uri[:len("http")] == "http" {
			ss := strings.Split(uri, "//")
			uri = ss[1]
		}

		// remove the '.git' extension.
		uri = uri[:len(uri)-len(path.Ext(uri))]

		// compare
		if uri != dir {
			fmt.Println("WARNING: " + uri + " does not match " + dir)
			fmt.Println("Do you want to continue anyway?")
			fmt.Println("Please type yes or no and then press enter:")
			if !askForConfirmation() {
				os.Exit(1)
			}
		}
	}

	return dir, nil
}
