package commands

import "github.com/spf13/cobra"

// GenerateCMD describes the GenerateCMD command.
var GenerateCMD = &cobra.Command{
	Use:   "generate",
	Short: "Generate boilerplate.",
	Long:  "Generate boilerplate for controllers, middleware, and models.",
	Run:   func(cmd *cobra.Command, args []string) {},
}
