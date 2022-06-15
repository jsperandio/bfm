package ui

import (
	"time"

	"github.com/jsperandio/bfm/app/domain/service"
	"github.com/jsperandio/bfm/app/ui/block"
	"github.com/jsperandio/bfm/app/ui/constant"
	"github.com/jsperandio/bfm/app/ui/model"
	"github.com/jsperandio/bfm/app/ui/widget"
	"github.com/rivo/tview"
)

type Screen struct {
	*tview.Application
	screenLayer    *tview.Pages
	mainMenu       block.Block
	newProjectMenu block.Block
	paramForm      block.Block
	menuPages      *tview.Pages
	textView       *widget.FileView
	layoutView     *widget.LayoutView
	viewPages      *tview.Pages
	progressBar    *widget.ProgressDialog
}

func NewScreen(mkr service.ProjectMaker) *Screen {

	scrn := Screen{
		Application: tview.NewApplication(),
		screenLayer: tview.NewPages(),
		menuPages:   tview.NewPages(),
		viewPages:   tview.NewPages(),
		progressBar: widget.NewProgressDialog(),
	}

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

	// Start Main Menu
	rfrs := model.NewRefers()
	rfrs.Add("menuPages", scrn.menuPages)

	mn := block.NewMainMenu(rfrs)
	scrn.mainMenu = mn

	// Start New Project Menu
	rfrs.Add("viewPages", scrn.viewPages)
	rfrs.Add("layoutView", scrn.layoutView)

	pm, err := block.NewProjectMenu(rfrs)
	if err != nil {
		panic(err)
	}
	scrn.newProjectMenu = pm

	// Start Param Form
	rfrs.Add("screenLayer", scrn.screenLayer)
	rfrs.Add("progressBar", scrn.progressBar)
	form := block.NewParamForm(rfrs, &model.Layout{}, mkr)
	scrn.paramForm = form

	// Add Menu Flow Pages
	scrn.menuPages.AddPage(scrn.mainMenu.GetName(), scrn.mainMenu, true, true)
	scrn.menuPages.AddPage(scrn.newProjectMenu.GetName(), scrn.newProjectMenu, true, false)
	scrn.menuPages.AddPage(scrn.paramForm.GetName(), form, true, false)

	// Add Screen Layer Pages
	flex := tview.NewFlex().
		AddItem(scrn.menuPages, 0, constant.ScreenMenuProportion, true).
		AddItem(scrn.viewPages, 0, constant.ScreenViewProportion, false)

	scrn.screenLayer = scrn.screenLayer.
		AddPage("default", flex, true, true).
		AddPage(constant.ModalProgress, scrn.progressBar, true, false)

	return &scrn
}

func (s *Screen) Render() error {

	go s.refreshChan()

	err := s.SetRoot(s.screenLayer, true).EnableMouse(false).Run()
	if err != nil {
		return err
	}

	return nil
}

func (s *Screen) refreshChan() {
	tick := time.NewTicker(constant.ScreenRefreshRate)
	for {
		select {
		case <-tick.C:

			s.Draw()
		}
	}
}
