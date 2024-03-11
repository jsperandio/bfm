package widget

import (
	"errors"
	"io/ioutil"

	"github.com/jsperandio/bfm/app/ui/constant"
	"github.com/rivo/tview"
)

type FileView struct {
	*tview.TextView
	name     string
	fileName string
	path     string
}

func NewFileView(fn string, pth string) (*FileView, error) {
	if fn == "" {
		return nil, errors.New("File name is empty")
	}

	if pth == "" {
		return nil, errors.New("Path is empty")
	}

	// open file from disk
	f, err := ioutil.ReadFile(pth + fn)
	if err != nil {
		return nil, err
	}

	fv := &FileView{
		TextView: tview.NewTextView(),
		fileName: fn,
		path:     pth,
		name:     constant.FileViewName,
	}

	fv.SetBorder(true).SetTitle(fv.fileName)
	fv.SetBorderPadding(1, 0, 2, 2)
	fv.SetText(string(f))

	return fv, nil
}

func (fv *FileView) GetName() string {
	return fv.name
}
