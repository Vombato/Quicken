package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "qk",
	Short: "Trying Cobra",
	Long: `Trying up Cobra for Quicken
				Quicken is a command line tool to generate software project based on clear, repeatable recipes`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
