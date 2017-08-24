package cmd

import (
	"github.com/spf13/cobra"
	"github.com/matteojoliveau/quicken/recipe"
	//"github.com/matteojoliveau/quicken/modules/git"
	"github.com/matteojoliveau/quicken/modules/readme"
	"fmt"
	"github.com/matteojoliveau/quicken/utils"
	"github.com/matteojoliveau/quicken/modules/git"
	"github.com/matteojoliveau/quicken/modules/license"
)

var recipePath string

func init() {
	initCmd.Flags().StringVarP(&recipePath, "recipe", "r", "recipe.yml", "Path to the recipe file")
	RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project",
	Long:  `Initialize a new project based on a recipe file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Quicken Generator v%s\n", utils.Version)
		r := recipe.ParseRecipe(recipePath)




		//Git
		gitR := r.Git
		if gitR.Init {
			git.InitRepo()

			rem := gitR.Remote
			if rem.Name != "" && rem.Url == "" {
				fmt.Println("Git Remote malformed configuration\n  Missing 'url' field in recipe file")
			} else if rem.Name == "" && rem.Url != "" {
				fmt.Println("Git Remote malformed configuration\n  Missing 'name' field in recipe file")
			} else if rem.Name != "" && rem.Url != "" {
				git.AddRemote(rem.Name, rem.Url)
			}
		}

		//README file
		if r.Readme {
			readme.Create(&readme.Readme{Title: r.Name, Headline: r.Description})
		}

		//LICENSE file
		if r.License != "" {
			license.Create(r.License, r.Owner.Name)
		}
	},
}
