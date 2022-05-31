package ui

import (
	"github.com/jsperandio/bfm/app/ui/block"
	"github.com/jsperandio/bfm/app/ui/model"
	"github.com/jsperandio/bfm/app/ui/widget"
	"github.com/rivo/tview"
)

type Screen struct {
	*tview.Application
	mainMenu       block.Menu
	newProjectMenu block.Menu
	paramForm      block.ParamForm
	menuPages      *tview.Pages
	textView       *widget.FileView
	layoutView     *widget.LayoutView
	viewPages      *tview.Pages
}

func NewScreen() *Screen {
	scrn := Screen{
		Application: tview.NewApplication(),
		menuPages:   tview.NewPages(),
		viewPages:   tview.NewPages(),
	}

	// Start Main Menu
	rfrs := model.NewRefers()
	rfrs.Add("menuPages", scrn.menuPages)

	mn := block.NewMainMenu(rfrs)
	scrn.mainMenu = mn

	// Start New Project Menu
	pm, err := block.NewProjectMenu(rfrs)
	if err != nil {
		panic(err)
	}
	scrn.newProjectMenu = pm

	// Start Param Form
	form := block.NewParamForm(&model.Layout{})
	form.StickyToPage(scrn.menuPages)
	scrn.paramForm = form

	// Add Menu Flow Pages
	scrn.menuPages.AddPage(scrn.mainMenu.GetName(), scrn.mainMenu, true, true)
	scrn.menuPages.AddPage(scrn.newProjectMenu.GetName(), scrn.newProjectMenu, true, false)
	scrn.menuPages.AddPage(scrn.paramForm.GetName(), form, true, false)

	// Start Viewer
	iv, err := widget.NewFileView("README.md", "./")
	if err != nil {
		panic(err)
	}
	scrn.textView = iv

	// Start Layout View
	lt := widget.NewLayoutView()
	scrn.layoutView = lt

	// Add View Flow Pages
	scrn.viewPages.AddPage(scrn.textView.GetName(), scrn.textView, true, true)
	scrn.viewPages.AddPage(scrn.layoutView.GetName(), scrn.layoutView, true, false)

	return &scrn
}

func (s *Screen) Render() error {

	flex := tview.NewFlex().
		AddItem(s.menuPages, 0, 3, true).
		AddItem(s.viewPages, 0, 4, false)

	err := s.SetRoot(flex, true).EnableMouse(false).Run()

	if err != nil {
		return err
	}

	return nil
}
