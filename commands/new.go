package commands

import (
	"log"

	"github.com/gophersaurus/gf/tool"
	"github.com/spf13/cobra"
)

var (
	newOrigin string
	newSkip   bool
)

func init() {
	NewCMD.Flags().StringVarP(&newOrigin, "origin", "o", "", "Create the project with a git remote origin")
	NewCMD.Flags().BoolVarP(&newSkip, "skip", "s", false, "Skip any warnings that would require user input")
}

// NewCMD describes the new command.
var NewCMD = &cobra.Command{
	Use:   "new",
	Short: "Update Gophersaurus.",
	Long:  "Update the Gophersaurus framework.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if err := tool.Create(args[0], newOrigin, verbose, newSkip, false); err != nil {
				log.Fatal(err)
			}
			return
		}

		if err := tool.Create("", newOrigin, verbose, newSkip, false); err != nil {
			log.Fatal(err)
		}
	},
}
