package blocks

import (
	"github.com/jsperandio/bfm/app/ui/model"
	"github.com/rivo/tview"
)

type Menu interface {
	tview.Primitive
	GetName() string
	StickyToPage(page *tview.Pages)
	UpdateItem(itemToUpdate string, item model.ListItem)
}
