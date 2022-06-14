package service

import (
	"errors"
	"fmt"
	"os"

	"github.com/jsperandio/bfm/app/domain/model"
	uimodel "github.com/jsperandio/bfm/app/ui/model"
)

type ProjectMaker interface {
	Make(layout *uimodel.Layout, prjct *uimodel.Project) error
}

type projectMaker struct{}

func NewProjectMaker() ProjectMaker {
	return &projectMaker{}
}

func (pm projectMaker) Make(lyt *uimodel.Layout, pjt *uimodel.Project) error {

	if lyt == nil {
		return errors.New("layout is nil")
	}

	if pjt == nil {
		return errors.New("project is nil")
	}

	project := model.NewProjectFromUI(pjt)
	layoutDir := model.NewLayoutDirFromUI(lyt)

	for _, p := range layoutDir.DirectPaths() {

		if fp := pm.buildPath(*project, p); fp != "" {

			err := os.MkdirAll(fp, os.ModePerm)
			if err != nil {
				return fmt.Errorf("failed to create directory %s: %s", fp, err)
			}
		}
	}
	return nil
}

func (pm projectMaker) buildPath(p model.Project, lytPth string) string {

	if p.RootPath == "" {
		return ""
	}

	// remove last slash if it exists
	if p.RootPath[len(p.RootPath)-1] == '/' {
		p.RootPath = p.RootPath[:len(p.RootPath)-1]
	}

	return fmt.Sprintf("%s/%s/%s", p.RootPath, p.Name, lytPth)
}
