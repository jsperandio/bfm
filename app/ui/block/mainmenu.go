package block

import (
	"github.com/jsperandio/bfm/app/ui/constant"
	"github.com/jsperandio/bfm/app/ui/model"
	"github.com/rivo/tview"
)

type mainMenu struct {
	*tview.List
	name       string
	items      map[string]*model.ListItem
	references *model.Refers
}

func NewMainMenu(r *model.Refers) Block {
	mm := &mainMenu{
		List:       tview.NewList(),
		name:       constant.MainMenuName,
		references: r,
	}

	mm.SetTitle(constant.MainMenuTitle).SetBorder(true)
	mm.SetBorderPadding(1, 0, 2, 0)

	mm.items = map[string]*model.ListItem{
		"New Project": {
			Index:       0,
			Text:        "New Project",
			Description: "Build a new project for you",
			Short:       'n',
			Selected: func() {
				page := mm.menuPages()
				if page != nil {
					page.SwitchToPage(constant.ProjectMenuName)
				}
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

func (mm *mainMenu) menuPages() *tview.Pages {
	return mm.references.Get("menuPages").AsPages()
}

func (mm *mainMenu) GetName() string {
	return mm.name
}

func (mm *mainMenu) SetRefers(r *model.Refers) {
	mm.references = r
}

func (mm *mainMenu) UpdateItem(itemToUpdate string, item model.ListItem) {
	rmvItem := mm.items[itemToUpdate]
	mm.RemoveItem(rmvItem.Index)
	mm.items[itemToUpdate] = &item
	mm.AddItem(item.Text, item.Description, item.Short, item.Selected)
}
