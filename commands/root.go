package commands

import "github.com/spf13/cobra"

var verbose bool

func init() {
	RootCMD.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Print verbose output")
}

// RootCMD describes the root command.
var RootCMD = &cobra.Command{
	Short: "The Gophersaurus project tool.",
	Long:  "The Gophersaurus CLI tool for Gophersaurus projects.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
