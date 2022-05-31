package block

import (
	"github.com/jsperandio/bfm/app/ui/model"
	"github.com/rivo/tview"
)

type Menu interface {
	tview.Primitive
	GetName() string
	SetRefers(r *model.Refers)
	// StickyToPage(page *tview.Pages)
	// UpdateItem(itemToUpdate string, item model.ListItem)
}
