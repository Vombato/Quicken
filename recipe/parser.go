package recipe

import (
	"gopkg.in/yaml.v2"
	"log"
	"fmt"
	"path/filepath"
	"io/ioutil"
)

type Recipe struct {
	Language   string
	Build_tool Tool `yaml:"build_tool"`
	Vcs        Tool `yaml:"vcs"`
	License    string
	Readme     bool
}

type Tool struct {
	Name    string
	Version string
}

func ParseRecipe(path string) {
	if !filepath.IsAbs(path) {
		abs, err := filepath.Abs(path)
		if err != nil {
			fmt.Println(err)
		}
		path = abs
	}
	data, ioerr := ioutil.ReadFile(path)
	if ioerr != nil {
		fmt.Println(ioerr)
	}
	recipe := Recipe{}
	err := yaml.Unmarshal([]byte(data), &recipe)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", recipe)
	fmt.Println(recipe.Language)
	fmt.Println(recipe.Vcs.Name)
	fmt.Println(recipe.Build_tool.Name)

}
