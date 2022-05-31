package widget

import (
	"errors"
	"io/ioutil"

	"github.com/rivo/tview"
)

const fileViewName = "FileView"

type FileView struct {
	*tview.TextView
	name     string
	fileName string
	path     string
}

func NewFileView(fn string, path string) (*FileView, error) {
	if fn == "" {
		return nil, errors.New("File name is empty")
	}

	if path == "" {
		return nil, errors.New("Path is empty")
	}

	// open file from disk
	f, err := ioutil.ReadFile(path + fn)
	if err != nil {
		return nil, err
	}

	fv := &FileView{
		TextView: tview.NewTextView(),
		fileName: fn,
		path:     path,
	}

	fv.SetBorder(true).SetTitle(fv.fileName)
	fv.SetBorderPadding(1, 0, 2, 0)
	fv.SetText(string(f))

	return fv, nil
}

func (fv *FileView) GetName() string {
	return fileViewName
}
