package block

import (
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/jsperandio/bfm/app/domain/service"
	"github.com/jsperandio/bfm/app/ui/constant"
	"github.com/jsperandio/bfm/app/ui/constant/label"
	"github.com/jsperandio/bfm/app/ui/converter"
	"github.com/jsperandio/bfm/app/ui/model"
	"github.com/jsperandio/bfm/app/ui/widget"
	"github.com/rivo/tview"
	"golang.org/x/sync/errgroup"
)

type paramForm struct {
	*tview.Form
	name         string
	layout       *model.Layout
	references   *model.Refers
	projectMaker service.ProjectMaker
}

func NewParamForm(rfs *model.Refers, lyt *model.Layout, pm service.ProjectMaker) Block {

	pf := &paramForm{
		Form:         tview.NewForm(),
		name:         constant.ParamFormName,
		layout:       lyt,
		references:   rfs,
		projectMaker: pm,
	}

	pf.newInitialPathField()
	pf.newProjectNameField()

	pf.newRunGoModField()

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

func (pf *paramForm) getFormItemValueByLabel(lb string) (value string) {

	fi := pf.GetFormItemByLabel(lb)

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

func (pf *paramForm) toggleGoMod(chk bool) {
	// remove git options if go mod is not run
	if !chk {
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

	pf.screenLayer().SendToFront(constant.ModalProgress)
	pf.screenLayer().ShowPage(constant.ModalProgress)
	pf.progressMaker().SetTitle("Creating Project...")
	pf.progressMaker().Display()

	np := &model.Project{
		RootPath:        pf.getFormItemValueByLabel(label.InputFieldInitialPath),
		GitPlatform:     pf.getFormItemValueByLabel(label.DropDownGitPlatform),
		GitUser:         pf.getFormItemValueByLabel(label.InputFieldGitUser),
		Name:            pf.getFormItemValueByLabel(label.InputFieldProjectName),
		RunGoMod:        pf.getFormItemValueByLabel(label.CheckboxRunGoMod),
		RememberChoices: pf.getFormItemValueByLabel(label.CheckboxRememberChoices),
	}

	errg := &errgroup.Group{}
	errg.Go(func() error {
		time.Sleep(time.Second * 1)
		err := pf.projectMaker.Make(pf.layout, np)
		if err != nil {
			return err
		}
		pf.progressMaker().Done()
		time.Sleep(time.Second * 1)
		return nil
	})

	go func() {
		errg.Wait()
		pf.screenLayer().HidePage(constant.ModalProgress)
		pf.screenLayer().SendToBack(constant.ModalProgress)
		pf.menuPages().SwitchToPage(constant.ProjectMenuName)
		return
	}()

}

func (pf *paramForm) menuPages() *tview.Pages {
	return converter.AsPages(pf.references.Get("menuPages"))
}

func (pf *paramForm) viewPages() *tview.Pages {
	return converter.AsPages(pf.references.Get("viewPages"))
}

func (pf *paramForm) screenLayer() *tview.Pages {
	return converter.AsPages(pf.references.Get("screenLayer"))
}

func (pf *paramForm) progressMaker() *widget.ProgressDialog {
	return converter.AsProgressDialog(pf.references.Get("progressBar"))
}
