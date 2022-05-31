package block

import (
	"io/ioutil"

	"github.com/jsperandio/bfm/app/ui/constant"
	"github.com/jsperandio/bfm/app/ui/converter"
	"github.com/jsperandio/bfm/app/ui/model"
	"github.com/jsperandio/bfm/app/ui/widget"
	"github.com/rivo/tview"
)

const ()

type ProjectMenu struct {
	*tview.List
	name           string
	items          map[string]*model.ListItem
	layouts        []*model.Layout
	selectedLayout int
	references     *model.Refers
}

func NewProjectMenu(r *model.Refers) (Block, error) {
	pm := &ProjectMenu{
		List:       tview.NewList(),
		name:       constant.ProjectMenuName,
		references: r,
	}

	pm.SetTitle(constant.ProjectMenuTitle).SetBorder(true)
	pm.SetBorderPadding(1, 0, 2, 0)

	err := pm.loadItems()
	if err != nil {
		return nil, err
	}

	for k, v := range pm.items {
		pm.AddItem(k, v.Description, v.Short, pm.selectLayoutEvent)
	}
	pm.addBackItem()

	return pm, nil
}

func (pm *ProjectMenu) loadItems() error {

	layouts, err := pm.getLayoutList()
	if err != nil {
		return err
	}

	pm.layouts = layouts
	pm.items = make(map[string]*model.ListItem)

	for i, l := range layouts {
		li := l.ToListItem()
		li.Short = rune(constant.KeyboardRunes[i])
		li.Selected = pm.selectLayoutEvent
		pm.items[l.FileName] = li
	}

	return nil
}

func (pm *ProjectMenu) addBackItem() {

	pm.items["Back"] = &model.ListItem{
		Index:       len(pm.items),
		Text:        "Back",
		Description: "Go back to the main menu",
		Short:       'b',
		Selected: func() {
			pm.menuPages().SwitchToPage(constant.MainMenuName)
		},
	}

	pm.AddItem(pm.items["Back"].Text, pm.items["Back"].Description, pm.items["Back"].Short, pm.items["Back"].Selected)
}

func (pm *ProjectMenu) selectLayoutEvent() {

	// can this go wrong???
	pm.selectedLayout = pm.GetCurrentItem()

	pm.menuPages().SwitchToPage(constant.ParamFormName)

	pm.viewPages().SwitchToPage(constant.LayoutViewName)

	pgname, fp := pm.menuPages().GetFrontPage()
	if pgname != constant.ParamFormName {
		return
	}

	pf, ok := fp.(*paramForm)
	if !ok {
		return
	}
	pf.SetLayout(pm.layouts[pm.selectedLayout])
	pm.layoutView().RenderLayout(pm.layouts[pm.selectedLayout])
}

func (pm *ProjectMenu) getLayoutList() ([]*model.Layout, error) {
	files, err := ioutil.ReadDir("./layouts")
	if err != nil {
		return nil, err
	}

	var layouts []*model.Layout
	for _, f := range files {
		layouts = append(layouts, model.NewLayout(f.Name()))
	}

	return layouts, nil
}

func (pm *ProjectMenu) menuPages() *tview.Pages {
	return converter.AsPages(pm.references.Get("menuPages"))
}

func (pm *ProjectMenu) viewPages() *tview.Pages {
	return converter.AsPages(pm.references.Get("viewPages"))
}

func (pm *ProjectMenu) layoutView() *widget.LayoutView {

	lv := pm.references.Get("layoutView")
	if lv == nil {
		return nil
	}

	return converter.AsLayoutView(lv)
}

func (pm *ProjectMenu) GetName() string {
	return pm.name
}

func (pm *ProjectMenu) SetRefers(r *model.Refers) {
	pm.references = r
}

func (pm *ProjectMenu) UpdateItem(itemToUpdate string, item model.ListItem) {
	rmvItem := pm.items[itemToUpdate]
	pm.RemoveItem(rmvItem.Index)
	pm.items[itemToUpdate] = &item
	pm.AddItem(item.Text, item.Description, item.Short, item.Selected)
}
