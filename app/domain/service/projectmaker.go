package service

import (
	"github.com/jsperandio/bfm/app/domain/model"
	uimodel "github.com/jsperandio/bfm/app/ui/model"
)

type ProjectMaker interface {
	MakeLayout(layout *uimodel.Layout) error
}

type projectMaker struct {
	project *model.Project
}

func NewProjectMaker(uiprjct uimodel.Project) ProjectMaker {
	return &projectMaker{
		project: model.NewProjectFromUI(&uiprjct),
	}
}

func (d *projectMaker) MakeLayout(layout *uimodel.Layout) error {

	return nil
}
