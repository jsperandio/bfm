package ui

import (
	"github.com/jsperandio/bfm/app/ui/blocks"
	"github.com/jsperandio/bfm/app/ui/model"
	"github.com/jsperandio/bfm/app/ui/widgets"
	"github.com/rivo/tview"
)

type Screen struct {
	*tview.Application
	mainMenu       blocks.Menu
	newProjectMenu blocks.Menu
	paramForm      blocks.ParamForm
	menuPages      *tview.Pages
	infoView       *widgets.FileView
}

func NewScreen() *Screen {
	scrn := Screen{
		Application: tview.NewApplication(),
		menuPages:   tview.NewPages(),
	}

	// Start Main Menu
	mn := blocks.NewMainMenu()
	scrn.mainMenu = mn

	// Start New Project Menu
	pm, err := blocks.NewProjectMenu()
	if err != nil {
		panic(err)
	}
	scrn.newProjectMenu = pm

	// Start Viewer
	iv, err := widgets.NewFileView("README.md", "./")
	if err != nil {
		panic(err)
	}
	scrn.infoView = iv

	// Start Param Form
	form := blocks.NewParamForm(model.Layout{})
	form.StickyToPage(scrn.menuPages)
	scrn.paramForm = form

	// Add Flow Pages
	scrn.addMenuPage(scrn.mainMenu.GetName(), scrn.mainMenu, true, false)
	scrn.addMenuPage(scrn.newProjectMenu.GetName(), scrn.newProjectMenu, true, false)
	scrn.menuPages.AddPage(scrn.paramForm.GetName(), form, true, true)

	return &scrn
}

func (s *Screen) addMenuPage(name string, menu blocks.Menu, resize bool, visible bool) {

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
