package modules

import (
	"github.com/matteojoliveau/quicken/modules/git"
	"github.com/matteojoliveau/quicken/modules/license"
)

type Module interface {
	GetName() string
	Execute(args []string) []error
}

func GetAllModules() map[string]Module {
	m := map[string]Module {
		"git": git.GitModule{},
		"license": license.License{},
	}
	return m
}
