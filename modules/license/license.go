package license

import (
	"time"
	"fmt"
	"log"
	"os"
	"github.com/matteojoliveau/quicken/utils"
	"text/template"
	"github.com/matteojoliveau/quicken/resources/templates/licenses"
	"github.com/matteojoliveau/quicken/modules"
)

type License struct {
	modules.Module
}
type licenseData struct {
	Year int
	Owner string
}

func (l *licenseData) GetName() string {
	return "license"
}

func (l *licenseData) Execute(args []string) []error {
	lic := args[0]
	name := args[1]

	err := create(lic, name)

	return []error{err}
}

func create(license string, name string) error {
	fmt.Println()
	year := time.Now().Year()
	l := &licenseData{Year: year, Owner: name}
	existent, err := utils.IsFileExistent("LICENSE")
	if err != nil {
		log.Fatal(err)
	} else {
		if existent {
			fmt.Println("LICENSE file already existent")
		} else {
			fmt.Println("Creating file 'LICENSE'")
			//abs, absErr := filepath.Abs(fmt.Sprintf("resources/templates/licenses/%s.tmpl", license))
			//if absErr != nil {
			//	return absErr
			//}
			//templateFile, tempErr := ioutil.ReadFile(abs)
			//if tempErr != nil {
			//	return tempErr
			//}

			templateFile := licenses.Licenses[license]

			t, parseErr := template.New("license").Parse(string(templateFile))
			if parseErr != nil {
				return parseErr
			}

			file, fErr := os.Create("LICENSE")
			if fErr != nil {
				return fErr
			}

			t.Execute(file, l)
		}
	}
}
