package blocks

import (
	"github.com/jsperandio/bfm/app/ui/model"
	"github.com/rivo/tview"
)

const (
	mainMenuTitle = "BFM - Build for Me - v1.0"
	mainMenuName  = "MainMenu"
)

type mainMenu struct {
	*tview.List
	name  string
	items map[string]*model.ListItem
	page  *tview.Pages
}

func NewMainMenu() Menu {
	mm := &mainMenu{
		List: tview.NewList(),
		name: mainMenuName,
	}

	mm.SetTitle(mainMenuTitle).SetBorder(true)
	mm.SetBorderPadding(1, 0, 2, 0)

	mm.items = map[string]*model.ListItem{
		"New Project": {
			Index:       0,
			Text:        "New Project",
			Description: "Build a new project for you",
			Short:       'n',
			Selected: func() {
				mm.page.SwitchToPage(projectMenuName)
			},
		},
		"Edit layouts": {
			Index:       1,
			Text:        "Edit layouts",
			Description: "Select a layout to edit",
			Short:       'e',
			Selected:    nil,
		},
	}

	for _, v := range mm.items {
		mm.AddItem(v.Text, v.Description, v.Short, v.Selected)
	}

	return mm
}

func (mm *mainMenu) GetName() string {
	return mm.name
}

func (mm *mainMenu) StickyToPage(page *tview.Pages) {
	mm.page = page
}

func (mm *mainMenu) UpdateItem(itemToUpdate string, item model.ListItem) {
	rmvItem := mm.items[itemToUpdate]
	mm.RemoveItem(rmvItem.Index)
	mm.items[itemToUpdate] = &item
	mm.AddItem(item.Text, item.Description, item.Short, item.Selected)
}
