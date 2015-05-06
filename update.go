package main

import "os/exec"

func update() error {

  // Update the Gophersaurus framework.
  if err := exec.Command("git", "pull", "upstream", "master").Run(); err != nil {
    return err
  }

  return nil
}
