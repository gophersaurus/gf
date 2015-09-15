package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string

func init() { version = "gf version 0.5.0" }

// VersionCMD describes the version command.
var VersionCMD = &cobra.Command{
	Use:   "version",
	Short: "Installed version of the gf tool.",
	Long:  "Print the current installed version of the gf tool.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
