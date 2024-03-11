package model

import (
	"strings"

	ui "github.com/jsperandio/bfm/app/ui/model"
)

type LayoutDir struct {
	fullpaths []string
}

func NewLayoutDirFromUI(lyt *ui.Layout) *LayoutDir {
	ld := &LayoutDir{}
	// string all maps to path notation in the file system
	// so we need to convert it to the path notation in the file system
	initPath, ok := lyt.Structure["dir"].(map[string]interface{})
	if !ok {
		return nil
	}
	endPaths := &[]string{}
	walkPath := &[]string{}

	ld.recPathBuild(walkPath, initPath, endPaths)
	ld.fullpaths = *endPaths

	return ld
}

func (ld *LayoutDir) recPathBuild(pth *[]string, vl interface{}, ep *[]string) {
	switch vl := vl.(type) {

	case map[string]interface{}:
		for k, v := range vl {
			*pth = append(*pth, k)
			ld.recPathBuild(pth, v, ep)
			*pth = (*pth)[:len(*pth)-1]
		}
	case nil:
		*ep = append(*ep, strings.Join(*pth, "/"))
	}
}

// Get the full paths of the dirs in layout.
//
//	Example: ["app/domain/service", "app/domain/service/mocks", ...]
func (ld *LayoutDir) DirectPaths() []string {
	return ld.fullpaths
}
