package model

import (
	"strings"

	ui "github.com/jsperandio/bfm/app/ui/model"
)

type LayoutDir struct {
	fullpaths []string
}

func NewLayoutDirFromUI(lyt *ui.Layout) *LayoutDir {

	//string all maps to path notation in the file system
	//so we need to convert it to the path notation in the file system

	initPath, ok := lyt.Structure["dir"].(map[string]interface{})
	if !ok {
		return nil
	}

	ld := &LayoutDir{}

	endpaths := &[]string{}
	fsPaths := &[]string{}

	ld.recPathBuild(fsPaths, initPath, endpaths)
	ld.fullpaths = *endpaths

	return ld
}

func (ld *LayoutDir) recPathBuild(paths *[]string, value interface{}, endPaths *[]string) {

	switch vl := value.(type) {

	case map[string]interface{}:
		for k, v := range vl {
			*paths = append(*paths, k)
			ld.recPathBuild(paths, v, endPaths)
			*paths = (*paths)[:len(*paths)-1]
		}
	case nil:
		*endPaths = append(*endPaths, strings.Join(*paths, "/"))
	}
}

func (ld *LayoutDir) DirectPaths() []string {
	return ld.fullpaths
}
