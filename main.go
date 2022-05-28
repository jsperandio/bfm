package main

import (
	"io/ioutil"

	"github.com/jsperandio/bfm/app/ui"
	"github.com/rivo/tview"
)

func getLayoutList() []string {
	files, err := ioutil.ReadDir("./layouts")
	if err != nil {
		panic(err)
	}

	var fileList []string
	for _, f := range files {
		fileList = append(fileList, f.Name())
	}

	return fileList

}

func InputInitialPath() *tview.InputField {
	initialPath := tview.NewInputField().
		SetLabel("Initial path:").
		SetFieldWidth(30).
		SetPlaceholder("/home/user/projects/")

	initialPath.SetBorder(true)

	return initialPath
}

func main() {

	mainScreen := ui.NewScreen()
	mainScreen.Render()

}
