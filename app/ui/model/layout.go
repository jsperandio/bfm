package model

import (
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v3"
)

type Layout struct {
	FileName  string
	FileExt   string
	Structure map[string]interface{}
}

func NewLayout(fileName string) *Layout {
	l := &Layout{}

	l.FileName, l.FileExt = l.explodeFileName(fileName)
	l.readStructure()
	return l
}

func (l *Layout) explodeFileName(fn string) (string, string) {
	if fn == "" {
		return "", ""
	}

	if !strings.Contains(fn, ".") {
		return fn, ""
	}

	strs := strings.Split(fn, ".")

	return strs[0], strs[1]
}

func (l *Layout) readStructure() {

	if l.FileExt != "yaml" {
		return
	}

	path := "layouts/" + l.GetFileNameWithExt()

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(file, &l.Structure)
	if err != nil {
		return
	}
}

func (l *Layout) ToListItem() *ListItem {
	return &ListItem{
		Text:        l.FileName,
		Description: l.GetFileNameWithExt(),
		Short:       rune(l.FileName[0]),
		Selected:    nil,
	}
}

func (l *Layout) GetFileNameWithExt() string {
	return l.FileName + "." + l.FileExt
}

func (l *Layout) Name() string {
	if l.Structure == nil {
		return ""
	}

	if name, ok := l.Structure["name"]; ok {
		return name.(string)
	}

	return ""
}
