package block

import (
	"github.com/jsperandio/bfm/app/ui/model"
	"github.com/rivo/tview"
)

type Block interface {
	tview.Primitive
	GetName() string
	SetRefers(r *model.Refers)
}
