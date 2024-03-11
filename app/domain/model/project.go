package model

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	ui "github.com/jsperandio/bfm/app/ui/model"
)

type Project struct {
	RootPath    string
	gitPlatform string
	gitUser     string
	Name        string
	runGoMod    bool
}

func NewProjectFromUI(uip *ui.Project) *Project {
	p := &Project{}
	p.setRootPathFromUI(uip.RootPath)
	p.setGitPlatformFromUI(uip.GitPlatform)
	p.setGitUserFromUI(uip.GitUser)
	p.setNameFromUI(uip.Name)
	p.setRunGoModFromUI(uip.RunGoMod)

	return p
}

func (p *Project) setRootPathFromUI(rootPath string) {
	p.RootPath = strings.TrimSpace(rootPath)
}

func (p *Project) setGitPlatformFromUI(gitPlatform string) {
	p.gitPlatform = strings.TrimSpace(gitPlatform)
}

func (p *Project) setGitUserFromUI(gitUser string) {
	p.gitUser = strings.TrimSpace(gitUser)
}

func (p *Project) setRunGoModFromUI(runGoMod string) {
	rgm, err := strconv.ParseBool(runGoMod)
	if err != nil {
		rgm = false
	}
	p.runGoMod = rgm
}

// If no name set generate one : awesome-project- + timestamp (seconds)
func (p *Project) setNameFromUI(name string) {
	if name == "" {
		p.Name = "awesome-project_" + time.Now().Format("1504")
	} else {
		p.Name = name
	}
}

// Get Full project name : gitplatform/gituser_name/name
func (p *Project) FullName() string {
	if p.gitUser == "" {
		return fmt.Sprintf("%s.com/%s", p.gitPlatform, p.Name)
	}

	return fmt.Sprintf("%s.com/%s/%s", p.gitPlatform, p.gitUser, p.Name)
}

func (p *Project) NeedRunGoMod() bool {
	return p.runGoMod
}
