package commands

import (
	"log"

	"github.com/gophersaurus/gf/generate"
	"github.com/spf13/cobra"
)

func init() {
	GenerateCMD.AddCommand(generateControllerCMD)
}

// GenerateCMD describes the GenerateCMD command.
var GenerateCMD = &cobra.Command{
	Use:   "generate",
	Short: "Generate boilerplate.",
	Long:  "Generate boilerplate for controllers, middleware, and models.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var generateControllerCMD = &cobra.Command{
	Use:   "controller",
	Short: "Generate a controller.",
	Long:  "Generate basic controller boilerplate.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := generate.Controller(args); err != nil {
			log.Fatal(err)
		}
	},
}
