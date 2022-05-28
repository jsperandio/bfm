package model

import "strings"

type Layout struct {
	FileName string
	FileExt  string
}

func NewLayout(fileName string) *Layout {
	l := &Layout{}

	l.FileName, l.FileExt = l.explodeFileName(fileName)
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
