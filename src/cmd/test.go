package cmd

import (
	"github.com/matteojoliveau/quicken/src/recipe"
	git "github.com/matteojoliveau/quicken/src/modules/git"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	//RootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Print the version number of Quicken",
	Long:  `All software has versions. This is Quicken's`,
	Run: func(cmd *cobra.Command, args []string) {
		parsedRecipe := recipe.ParseRecipe("test.yml")
		fmt.Printf("--- t:\n%v\n\n", parsedRecipe)
		err := git.InitRepo()
		if err != nil {
			log.Fatal(err)
		}
	},
}
