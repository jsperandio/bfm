package model

import (
	"strconv"
	"time"

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

	p := &Project{
		RootPath:    uip.RootPath,
		GitPlatform: uip.GitPlatform,
		GitUser:     uip.GitUser,
		RunGoMod:    rgm,
	}

	p.SetName(uip.Name)

	return p
}

// set name , if no name set generate one : awesome-project- + timestamp (seconds)
func (p *Project) SetName(name string) {
	if name == "" {
		p.Name = "awesome-project_" + time.Now().Format("1504")
	} else {
		p.Name = name
	}
}
