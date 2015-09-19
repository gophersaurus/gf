package commands

import (
	"log"

	"github.com/gophersaurus/gf/project"
	"github.com/spf13/cobra"
)

// UpdateCMD describes the new command.
var UpdateCMD = &cobra.Command{
	Use:   "update",
	Short: "Update the Gophersaurus framework.",
	Long:  "Update the Gophersaurus framework.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := project.Update(); err != nil {
			log.Fatal(err)
		}
	},
}
