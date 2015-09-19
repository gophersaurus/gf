package main

import "github.com/gophersaurus/gf/commands"

func main() {
	commands.RootCMD.AddCommand(commands.NewCMD)
	commands.RootCMD.AddCommand(commands.UpdateCMD)
	commands.RootCMD.AddCommand(commands.VersionCMD)
	commands.RootCMD.AddCommand(commands.GenerateCMD)
	commands.RootCMD.Execute()
}
