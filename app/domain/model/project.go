package model

import (
	"strconv"

	ui "github.com/jsperandio/bfm/app/ui/model"
)

type Project struct {
	RootPath    string
	GitPlatform string
	GitUser     string
	Name        string
	RunGoMod    bool
}

func NewProjectFromUI(uip *ui.Project) *Project {

	rgm, err := strconv.ParseBool(uip.RunGoMod)
	if err != nil {
		rgm = false
	}

	return &Project{
		RootPath:    uip.RootPath,
		GitPlatform: uip.GitPlatform,
		GitUser:     uip.GitUser,
		Name:        uip.Name,
		RunGoMod:    rgm,
	}
}
