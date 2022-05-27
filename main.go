package main

import (
	"io/ioutil"

	"github.com/rivo/tview"
)

var (
	// App is the tview application.
	app *tview.Application
	// pages is the main view of the application.
	pages *tview.Pages
)

func ListActions() *tview.List {
	startNewProjectPage := func() {
		pages.SwitchToPage("New Project")
	}

	listActions := tview.NewList()
	listActions.SetBorder(true).SetTitle("BFM - Build for Me - v1.0")
	listActions.AddItem("New Project", "Build a new project for you", 'n', startNewProjectPage)
	listActions.AddItem("Edit layouts", "Select a layout to edit", 'e', nil)

	return listActions
}

func NewProjectActions() *tview.List {
	returnToMenu := func() {
		pages.SwitchToPage("Actions")
	}

	getFileList := func() []string {
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

	keyboardShortcuts := []rune{'q', 'w', 'e', 'r', 'a', 's', 'd'}

	listActions := tview.NewList()
	listActions.SetBorder(true).SetTitle("New Project")

	files := getFileList()
	for i, f := range files {

		listActions.AddItem(f, "", keyboardShortcuts[i], nil)
	}

	// listActions.AddItem("Hexame", "Hexame.yaml", 'h', nil)
	// listActions.AddItem("Clean", "Clean.yaml", 'c', nil)
	listActions.AddItem("Back", "Return to menu", 'b', returnToMenu)
	return listActions
}

func Readme() *tview.TextView {
	readme := tview.NewTextView()
	readme.SetBorder(true).SetTitle("Readme")

	// open readme.md file from disk
	f, err := ioutil.ReadFile("README.md")
	if err != nil {
		panic(err)
	}

	// set readme.md content
	readme.SetText(string(f))

	return readme
}

func main() {
	// editor := tview.NewTable().SetBorders(true)
	// editor.SetBorder(true).SetTitle("Editor")

	pages = tview.NewPages()
	pages.AddPage("Actions", ListActions(), true, true)
	pages.AddPage("New Project", NewProjectActions(), true, true)

	pages.SwitchToPage("Actions")

	// Create the layout.
	flex := tview.NewFlex().
		AddItem(pages, 0, 1, true).
		AddItem(Readme(), 0, 5, false)
		// SetTitle("BFM - Build for Me - v1.0")

	app = tview.NewApplication()

	app.SetRoot(flex, true)
	app.Run()

}
