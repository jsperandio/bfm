package model

import "strings"

type Layout struct {
	FileName string
	FileExt  string
}

func NewLayout(fileName string) *Layout {
	l := &Layout{
		FileName: fileName,
	}

	l.FileExt = l.getExt()

	return l
}

func (l *Layout) getExt() string {

	if strings.Contains(l.FileName, ".") {
		return strings.Split(l.FileName, ".")[1]
	}

	return ""
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
