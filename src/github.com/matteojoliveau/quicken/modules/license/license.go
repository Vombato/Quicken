package license

import (
	"time"
	"fmt"
	"log"
	"os"
	"github.com/matteojoliveau/quicken/utils"
	"text/template"
	"github.com/matteojoliveau/quicken/resources/templates/licenses"
)

type License struct {
	Year int
	Owner string
}

func Create(license string, name string) {
	fmt.Println()
	year := time.Now().Year()
	l := &License{Year: year, Owner: name}
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
			//	log.Fatal(absErr)
			//}
			//templateFile, tempErr := ioutil.ReadFile(abs)
			//if tempErr != nil {
			//	log.Fatal(tempErr)
			//}

			templateFile := licenses.Licenses[license]

			t, parseErr := template.New("license").Parse(string(templateFile))
			if parseErr != nil {
				log.Fatal(parseErr)
			}

			file, fErr := os.Create("LICENSE")
			if fErr != nil {
				log.Fatal(fErr)
			}

			t.Execute(file, l)
		}
	}
}
