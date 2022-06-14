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

type projectMaker struct {
	goModules GoInitializer
}

func NewProjectMaker(gm GoInitializer) ProjectMaker {
	return &projectMaker{
		goModules: gm,
	}
}

func (pm projectMaker) Make(lyt *uimodel.Layout, pjct *uimodel.Project) error {

	if lyt == nil {
		return errors.New("layout is nil")
	}

	if pjct == nil {
		return errors.New("project is nil")
	}

	project := model.NewProjectFromUI(pjct)
	layoutDir := model.NewLayoutDirFromUI(lyt)

	err := pm.buildDirectoryTree(layoutDir, project)
	if err != nil {
		return err
	}

	if project.NeedRunGoMod() {

		err = pm.goModules.Init(project)
		if err != nil {
			return err
		}
	}

	return nil
}

func (pm projectMaker) buildDirectoryTree(layoutDir *model.LayoutDir, project *model.Project) error {

	for _, path := range layoutDir.DirectPaths() {

		if pp := pm.buildPathForProject(*project, path); pp != "" {

			err := os.MkdirAll(pp, os.ModePerm)
			if err != nil {
				return fmt.Errorf("failed to create directory %s: %s", pp, err)
			}
		}
	}
	return nil
}

func (pm projectMaker) buildPathForProject(pjct model.Project, path string) string {

	if pjct.RootPath == "" {
		return ""
	}

	// remove last slash if it exists
	if pjct.RootPath[len(pjct.RootPath)-1] == '/' {
		pjct.RootPath = pjct.RootPath[:len(pjct.RootPath)-1]
	}

	return fmt.Sprintf("%s/%s/%s", pjct.RootPath, pjct.Name, path)
}
