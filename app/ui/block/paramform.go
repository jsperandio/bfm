package block

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jsperandio/bfm/app/ui/constant"
	"github.com/jsperandio/bfm/app/ui/model"
	"github.com/rivo/tview"
)

type paramForm struct {
	*tview.Form
	name       string
	layout     *model.Layout
	references *model.Refers
}

func NewParamForm(r *model.Refers, sl *model.Layout) Block {

	pf := &paramForm{
		Form:       tview.NewForm(),
		name:       constant.ParamFormName,
		layout:     sl,
		references: r,
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

func (pf *paramForm) GetName() string {
	return pf.name
}

func (pf *paramForm) SetRefers(r *model.Refers) {
	pf.references = r
}

func (pf *paramForm) SetLayout(l *model.Layout) {
	pf.layout = l
	pf.updateTitle()
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
	pf.menuPages().SwitchToPage(constant.ProjectMenuName)
}

func (pf *paramForm) menuPages() *tview.Pages {
	return pf.references.Get("menuPages").AsPages()
}
