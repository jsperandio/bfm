package model

import (
	ui "github.com/jsperandio/bfm/app/ui/model"
)

type LayoutDir struct {
	fullpaths []string
}

func NewLayoutDirFromUI(lyt *ui.Layout) *LayoutDir {

	//string all maps to path notation in the file system
	//so we need to convert it to the path notation in the file system
	fsPaths := make([]string, 0)

	initPath, ok := lyt.Structure["dir"].(map[string]interface{})
	if !ok {
		return nil
	}

	var stack []interface{}
	var pathStack []string

	stack = append(stack, initPath)

	for len(stack) > 0 {
		top := stack[len(stack)-1].(map[string]interface{})
		stack = stack[:len(stack)-1]

		last := ""
		for k, v := range top {

			if v == nil {
				// remove from pathStack
				top := pathStack[len(pathStack)-1]
				fsPaths = append(fsPaths, top+"/"+k)
				continue
			}

			if _, ok := v.(map[string]interface{}); ok {
				stack = append(stack, v)
			}

			last = k
		}

		if len(pathStack) == 0 {
			pathStack = append(pathStack, last)
			continue
		}

		pathStack = append(pathStack, pathStack[len(pathStack)-1]+"/"+last)
	}

	return &LayoutDir{
		fullpaths: fsPaths,
	}
}

func (ld *LayoutDir) DirectPaths() []string {
	return ld.fullpaths
}
