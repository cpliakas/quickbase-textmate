package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quickbase-textmate",
	Short: "Generates aspects of TextMate grammar for Quickbase formulae.",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute runs the command line tool.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
