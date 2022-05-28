package blocks

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

type projectMenu struct {
	*tview.List
	page           *tview.Pages
	name           string
	items          map[string]*model.ListItem
	layouts        []model.Layout
	selectedLayout int
}

func NewProjectMenu() (Menu, error) {
	pm := &projectMenu{
		List: tview.NewList(),
		name: projectMenuName,
	}

	pm.SetTitle(projectMenuTitle).SetBorder(true)

	err := pm.loadItems()
	if err != nil {
		return nil, err
	}

	for k, v := range pm.items {
		pm.AddItem(k, v.Description, v.Short, pm.selectLayout)
	}
	pm.addBackItem()

	return pm, nil
}

func (pm *projectMenu) loadItems() error {

	layouts, err := pm.getLayoutList()
	if err != nil {
		return err
	}

	pm.layouts = layouts
	pm.items = make(map[string]*model.ListItem)

	for i, l := range layouts {
		li := l.ToListItem()
		li.Short = rune(keyboardRunes[i])
		li.Selected = pm.selectLayout
		pm.items[l.FileName] = li
	}

	return nil
}

func (pm *projectMenu) addBackItem() {

	pm.items["Back"] = &model.ListItem{
		Index:       len(pm.items),
		Text:        "Back",
		Description: "Go back to the main menu",
		Short:       'b',
		Selected: func() {
			pm.page.SwitchToPage(mainMenuName)
		},
	}

	pm.AddItem(pm.items["Back"].Text, pm.items["Back"].Description, pm.items["Back"].Short, pm.items["Back"].Selected)
}

func (pm *projectMenu) selectLayout() {

	// can this go wrong???

	pm.selectedLayout = pm.GetCurrentItem()

	pm.page.SwitchToPage(paramFormName)

	pgname, fp := pm.page.GetFrontPage()
	if pgname != paramFormName {
		return
	}

	pf, ok := fp.(*paramForm)
	if !ok {
		return
	}

	pf.SetLayout(pm.layouts[pm.selectedLayout])

}

func (pm *projectMenu) getLayoutList() ([]model.Layout, error) {
	files, err := ioutil.ReadDir("./layouts")
	if err != nil {
		return nil, err
	}

	var layouts []model.Layout
	for _, f := range files {
		layouts = append(layouts, *model.NewLayout(f.Name()))
	}

	return layouts, nil
}

func (pm *projectMenu) GetName() string {
	return pm.name
}

func (pm *projectMenu) StickyToPage(page *tview.Pages) {
	pm.page = page
}

func (pm *projectMenu) UpdateItem(itemToUpdate string, item model.ListItem) {
	rmvItem := pm.items[itemToUpdate]
	pm.RemoveItem(rmvItem.Index)
	pm.items[itemToUpdate] = &item
	pm.AddItem(item.Text, item.Description, item.Short, item.Selected)
}
