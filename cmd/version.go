package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Quicken",
	Long:  `All software has versions. This is Quicken's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`Quicken Project Generator v0.1`)
	},
}
