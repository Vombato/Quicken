package recipe

import (
	"gopkg.in/yaml.v2"
	"log"
	"path/filepath"
	"io/ioutil"
)

type Recipe struct {
	Name        string
	Description string
	Owner struct {
		Name string
		Email string
		Organization string
		Url string
	}
	Language    string
	Build_tool  Tool `yaml:"build_tool"`
	Git         Git `yaml:"git"`
	License     string
	Readme      bool
}

type Git struct {
	Init bool
	Remote struct {
		Name string
		Url  string
	}
}

type Tool struct {
	Name    string
	Version string
}

func ParseRecipe(path string) Recipe {
	if !filepath.IsAbs(path) {
		abs, err := filepath.Abs(path)
		if err != nil {
			log.Fatal(err)
		}
		path = abs
	}
	data, ioerr := ioutil.ReadFile(path)
	if ioerr != nil {
		log.Fatal(ioerr)
	}
	recipe := Recipe{}
	err := yaml.Unmarshal([]byte(data), &recipe)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//fmt.Printf("--- t:\n%v\n\n", recipe)
	//fmt.Println(recipe.Language)
	//fmt.Println(recipe.Vcs.Name)
	//fmt.Println(recipe.Build_tool.Name)
	return recipe
}
