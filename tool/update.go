package tool

import "os/exec"

func Update() error {

	// Update the Gophersaurus framework.
	if err := exec.Command("git", "pull", "upstream", "master").Run(); err != nil {
		return err
	}

	return nil
}
