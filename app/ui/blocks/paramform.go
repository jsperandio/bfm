package blocks

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jsperandio/bfm/app/ui/model"
	"github.com/rivo/tview"
)

const paramFormName = "ParamForm"

type paramForm struct {
	*tview.Form
	name   string
	pages  *tview.Pages
	layout *model.Layout
}

type ParamForm interface {
	tview.Primitive
	StickyToPage(page *tview.Pages)
	GetName() string
}

func NewParamForm(sl *model.Layout) ParamForm {

	pf := &paramForm{
		Form:   tview.NewForm(),
		name:   paramFormName,
		layout: sl,
	}

	pf.AddInputField("Initial path", "/documents/repository", 30, nil, nil).
		AddDropDown("Git plataform", []string{" gitlab ", " github "}, 0, nil).
		AddInputField("Git User", "johndoe", 15, nil, nil).
		AddInputField("Project Name", "awesomeproject", 30, nil, nil).
		AddCheckbox("Run go mod", true, nil).
		AddCheckbox("Remember choices", true, nil).
		AddButton("Start", nil).
		AddButton("Cancel", pf.cancelAction)

	pf.customStyles()

	return pf
}

func (pf *paramForm) StickyToPage(page *tview.Pages) {
	pf.pages = page
}

func (pf *paramForm) GetName() string {
	return pf.name
}

func (pf *paramForm) customStyles() {

	pf.SetBorder(true)
	pf.updateTitle()
	pf.SetFieldBackgroundColor(tcell.Color53)
	pf.SetFieldTextColor(tcell.ColorWhite)
	pf.SetButtonBackgroundColor(tcell.ColorDarkBlue)
}

func (pf *paramForm) updateTitle() {
	pf.SetTitle("Start " + pf.layout.FileName + " Layout")
}

func (pf *paramForm) cancelAction() {
	pf.pages.SwitchToPage(projectMenuName)
}

func (pf *paramForm) SetLayout(l *model.Layout) {
	pf.layout = l
	pf.updateTitle()
}
