package block

import (
	"io/ioutil"

	"github.com/jsperandio/bfm/app/ui/model"
	"github.com/rivo/tview"
)

const (
	projectMenuTitle = "New Project"
	projectMenuName  = "NewProjectMenu"
	// removed [B] for back option
	keyboardRunes = "qwertyuiopasdfghjklzxcvnm"
)

type ProjectMenu struct {
	*tview.List
	name           string
	items          map[string]*model.ListItem
	layouts        []*model.Layout
	selectedLayout int
	references     *model.Refers
}

func NewProjectMenu(r *model.Refers) (Menu, error) {
	pm := &ProjectMenu{
		List:       tview.NewList(),
		name:       projectMenuName,
		references: r,
	}

	pm.SetTitle(projectMenuTitle).SetBorder(true)
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
		li.Short = rune(keyboardRunes[i])
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
			pm.menuPages().SwitchToPage(mainMenuName)
		},
	}

	pm.AddItem(pm.items["Back"].Text, pm.items["Back"].Description, pm.items["Back"].Short, pm.items["Back"].Selected)
}

func (pm *ProjectMenu) selectLayoutEvent() {

	// can this go wrong???
	pm.selectedLayout = pm.GetCurrentItem()

	pm.menuPages().SwitchToPage(paramFormName)
	pgname, fp := pm.menuPages().GetFrontPage()
	if pgname != paramFormName {
		return
	}

	pf, ok := fp.(*paramForm)
	if !ok {
		return
	}

	pf.SetLayout(pm.layouts[pm.selectedLayout])
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
	return pm.references.Get("menuPages").AsPages()
}

func (pm *ProjectMenu) GetName() string {
	return pm.name
}

func (pm *ProjectMenu) SetRefers(r *model.Refers) {
	pm.references = r
}

func (pm *ProjectMenu) StickyToPage(page *tview.Pages) {
}

func (pm *ProjectMenu) UpdateItem(itemToUpdate string, item model.ListItem) {
	rmvItem := pm.items[itemToUpdate]
	pm.RemoveItem(rmvItem.Index)
	pm.items[itemToUpdate] = &item
	pm.AddItem(item.Text, item.Description, item.Short, item.Selected)
}
