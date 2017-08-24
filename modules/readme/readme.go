package readme

import (
	"log"
	"text/template"
	"os"
	"fmt"
	"github.com/matteojoliveau/quicken/utils"
	"github.com/matteojoliveau/quicken/resources/templates"
)

type Readme struct {
	Title    string
	Headline string
}

func Create(r *Readme) {
	fmt.Println()
	existent, err := utils.IsFileExistent("README.md")
	if err != nil {
		log.Fatal(err)
	} else {
		if existent {
			fmt.Println("README file already existent")
		} else {
			fmt.Println("Creating file 'README.md'")
			//abs, absErr := filepath.Abs("resources/templates/Readme.md.tmpl")
			//if absErr != nil {
			//	log.Fatal(absErr)
			//}
			//templateFile, tempErr := ioutil.ReadFile(abs)
			//if tempErr != nil {
			//	log.Fatal(tempErr)
			//}

			templateFile := templates.Readme
			t, parseErr := template.New("readme").Parse(string(templateFile))
			if parseErr != nil {
				log.Fatal(parseErr)
			}

			file, fErr := os.Create("README.md")
			if fErr != nil {
				log.Fatal(fErr)
			}

			t.Execute(file, r)
		}
	}

}
