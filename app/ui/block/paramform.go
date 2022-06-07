package block

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/jsperandio/bfm/app/domain/service"
	"github.com/jsperandio/bfm/app/ui/constant"
	"github.com/jsperandio/bfm/app/ui/constant/label"
	"github.com/jsperandio/bfm/app/ui/converter"
	"github.com/jsperandio/bfm/app/ui/model"
	"github.com/rivo/tview"
)

type paramForm struct {
	*tview.Form
	name       string
	layout     *model.Layout
	references *model.Refers
	maker      service.ProjectMaker
}

func NewParamForm(r *model.Refers, l *model.Layout, pm service.ProjectMaker) Block {

	pf := &paramForm{
		Form:       tview.NewForm(),
		name:       constant.ParamFormName,
		layout:     l,
		references: r,
		maker:      pm,
	}

	pf.newInitialPathField()
	pf.newProjectNameField()

	pf.newRunGoModField()
	// pf.newGitPlatformField()
	// pf.newGitUserField()

	pf.newRememberChoicesField()
	pf.newStartButton()
	pf.newCancelButton()

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

func (pf *paramForm) newInitialPathField() {
	pf.AddInputField(label.InputFieldInitialPath, "./documents/repository", 30, nil, nil)
}

func (pf *paramForm) newGitPlatformField() {
	pf.AddDropDown(label.DropDownGitPlatform, []string{" gitlab ", " github "}, 0, nil)
}

func (pf *paramForm) newGitUserField() {
	pf.AddInputField(label.InputFieldGitUser, "johndoe", 0, nil, nil)
}

func (pf *paramForm) newProjectNameField() {
	pf.AddInputField(label.InputFieldProjectName, "", 30, nil, nil)
}

func (pf *paramForm) newRunGoModField() {
	pf.AddCheckbox(label.CheckboxRunGoMod, false, pf.toggleGoMod)
}

func (pf *paramForm) newRememberChoicesField() {
	pf.AddCheckbox(label.CheckboxRememberChoices, true, nil)
}

func (pf *paramForm) newStartButton() {
	pf.AddButton(label.ButtonStart, pf.startAction)
}

func (pf *paramForm) newCancelButton() {
	pf.AddButton(label.ButtonCancel, pf.cancelAction)
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

func (pf *paramForm) getFormItemValueByLabel(label string) (value string) {

	fi := pf.GetFormItemByLabel(label)

	switch fi.(type) {

	case *tview.InputField:
		inpt, ok := fi.(*tview.InputField)
		if ok {
			value = inpt.GetText()
		}
	case *tview.DropDown:
		dd, ok := fi.(*tview.DropDown)
		if ok {
			_, value = dd.GetCurrentOption()
		}

	case *tview.Checkbox:
		cb, ok := fi.(*tview.Checkbox)
		if ok {
			value = strconv.FormatBool(cb.IsChecked())
		}

	}

	return value
}

func (pf *paramForm) toggleGoMod(checked bool) {
	// remove git options if go mod is not run
	if !checked {
		pf.RemoveFormItem(pf.GetFormItemIndex(label.InputFieldGitUser))
		pf.RemoveFormItem(pf.GetFormItemIndex(label.DropDownGitPlatform))
	} else {
		pf.newGitPlatformField()
		pf.newGitUserField()
	}

	pf.RemoveFormItem(pf.GetFormItemIndex(label.CheckboxRememberChoices))
	pf.newRememberChoicesField()
}
func (pf *paramForm) cancelAction() {
	pf.SetFocus(0)
	pf.menuPages().SwitchToPage(constant.ProjectMenuName)
	pf.viewPages().SwitchToPage(constant.FileViewName)
}

func (pf *paramForm) startAction() {

	np := &model.Project{
		RootPath:        pf.getFormItemValueByLabel(label.InputFieldInitialPath),
		GitPlatform:     pf.getFormItemValueByLabel(label.DropDownGitPlatform),
		GitUser:         pf.getFormItemValueByLabel(label.InputFieldGitUser),
		Name:            pf.getFormItemValueByLabel(label.InputFieldProjectName),
		RunGoMod:        pf.getFormItemValueByLabel(label.CheckboxRunGoMod),
		RememberChoices: pf.getFormItemValueByLabel(label.CheckboxRememberChoices),
	}

	err := pf.maker.Make(pf.layout, np)
	if err != nil {
		return
	}
}

func (pf *paramForm) menuPages() *tview.Pages {
	return converter.AsPages(pf.references.Get("menuPages"))
}

func (pf *paramForm) viewPages() *tview.Pages {
	return converter.AsPages(pf.references.Get("viewPages"))
}
