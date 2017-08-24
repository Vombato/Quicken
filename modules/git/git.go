package git

import (
	"os/exec"
	"fmt"
	"github.com/matteojoliveau/quicken/utils"
	"github.com/matteojoliveau/quicken/modules"
)

type GitModule struct {
	modules.Module
}

func (g *GitModule) Execute(args []string) []error {
	var e []error
	for i := 0; i < len(args); i++ {
		c := args[i]
		switch c {
		case "init":
			err := initRepo()
			e = append(e, err)
		case "remote":
			err := addRemote(args[i+1], args[i+2])
			e = append(e, err)
		}
	}

	return e
}

func (g *GitModule) GetName() string {
	return "git"
}

func initRepo() error {
	fmt.Println()
	existent, err := utils.IsFileExistent(".git/")
	if err != nil {
		return err
	} else {
		if existent {
			fmt.Println("Git repository already existent")
			return nil
		} else {
			fmt.Println("Initializing Git repository")
			cmd := exec.Command("git", "init")
			output, err := cmd.CombinedOutput()
			fmt.Printf("%s\n", output)
			return err
		}
	}
}

func addRemote(name string, url string) error {
	fmt.Printf("\nAdding remote repository '%s' with URL %s\n", name, url)

	cmd := exec.Command("git", "remote", "add", name, url)
	output, err := cmd.CombinedOutput()
	fmt.Printf("%s\n", output)
	return err
}
