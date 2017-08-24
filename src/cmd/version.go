package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/matteojoliveau/quicken/src/utils"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Quicken",
	Long:  `All software has versions. This is Quicken's`,
	Run: func(cmd *cobra.Command, args []string) {
		ver := utils.Version
		fmt.Printf("Quicken Project Generator\n Version: %s\n", ver)
	},

}
