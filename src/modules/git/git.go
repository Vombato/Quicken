package git

import (
	"os/exec"
	"fmt"
	"github.com/matteojoliveau/quicken/src/utils"
)

func InitRepo() error {
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

func AddRemote(name string, url string) error {
	fmt.Printf("\nAdding remote repository '%s' with URL %s\n", name, url)

	cmd := exec.Command("git", "remote", "add", name, url)
	output, err := cmd.CombinedOutput()
	fmt.Printf("%s\n", output)
	return err
}
