package service

import (
	"os/exec"

	"github.com/jsperandio/bfm/app/domain/model"
)

// this is the service that initializes the golang module with the project

type GoInitializer interface {
	Init(p *model.Project) error
}

type goInitializer struct{}

func NewGoInitializer() GoInitializer {
	return &goInitializer{}
}

// The go mod init command initializes and writes a new go.mod file in the current directory,
// in effect creating a new module rooted at the current directory.
// 	The go.mod file must not already exist.
func (*goInitializer) Init(p *model.Project) error {

	app := "go"
	args := []string{"mod", "init", p.FullName()}

	cmd := exec.Command(app, args...)
	cmd.Dir = p.RootPath
	_, err := cmd.Output()

	if err != nil {
		return err
	}
	return nil
}
