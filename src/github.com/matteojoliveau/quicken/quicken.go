package quicken

import (
	"fmt"
	"os"

	"github.com/matteojoliveau/quicken/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
