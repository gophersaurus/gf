package commands

import (
	"log"

	"github.com/gophersaurus/gf/project"
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
	Short: "Create a new project.",
	Long:  "Create a new gophersaurus project.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if err := project.Create(args[0], newOrigin, verbose, newSkip, false); err != nil {
				log.Fatal(err)
			}
			return
		}

		if err := project.Create("", newOrigin, verbose, newSkip, false); err != nil {
			log.Fatal(err)
		}
	},
}
