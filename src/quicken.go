package src

import (
	"fmt"
	"os"

	"github.com/matteojoliveau/quicken/src/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
