package ui

import (
	"github.com/jsperandio/bfm/app/ui/blocks"
	"github.com/jsperandio/bfm/app/ui/widgets"
	"github.com/rivo/tview"
)

type Screen struct {
	*tview.Application
	mainMenu       blocks.Menu
	newProjectMenu blocks.Menu
	menuPages      *tview.Pages
	infoView       *widgets.FileView
}

func NewScreen() *Screen {
	scrn := Screen{
		Application: tview.NewApplication(),
		menuPages:   tview.NewPages(),
	}

	mn := blocks.NewMainMenu()
	scrn.mainMenu = mn

	pm, err := blocks.NewProjectMenu()
	if err != nil {
		panic(err)
	}
	scrn.newProjectMenu = pm

	iv, err := widgets.NewFileView("README.md", "./")
	if err != nil {
		panic(err)
	}
	scrn.infoView = iv

	scrn.AddMenuPage(scrn.mainMenu.GetName(), scrn.mainMenu, true, true)
	scrn.AddMenuPage(scrn.newProjectMenu.GetName(), scrn.newProjectMenu, true, false)

	// app.menuPages.AddPage("InputInitialPath", InputInitialPath(), true, false)

	return &scrn
}

func (s *Screen) AddMenuPage(name string, menu blocks.Menu, resize bool, visible bool) {

	menu.StickyToPage(s.menuPages)
	s.menuPages.AddPage(name, menu, resize, visible)

}

func (s *Screen) Render() error {

	flex := tview.NewFlex().
		AddItem(s.menuPages, 0, 3, true).
		AddItem(s.infoView, 0, 4, false)

	err := s.SetRoot(flex, true).EnableMouse(false).Run()

	if err != nil {
		return err
	}

	return nil
}
