package converter

import (
	"log"

	"github.com/jsperandio/bfm/app/ui/widget"
	"github.com/rivo/tview"
)

func AsPages(value interface{}) *tview.Pages {

	pg, ok := value.(*tview.Pages)
	if ok {
		return pg
	}

	log.Println("[WARN] converter.AsPages: value is not a tview.Pages")
	return nil
}

func AsTreeView(value interface{}) *tview.TreeView {

	tv, ok := value.(*tview.TreeView)
	if ok {
		return tv
	}

	log.Print("[WARN] converter.AsTreeView: value is not a tview.TreeView")
	return nil
}

func AsLayoutView(value interface{}) *widget.LayoutView {

	lt, ok := value.(*widget.LayoutView)
	if ok {
		return lt
	}
	log.Print("[WARN] converter.AsLayoutView: value is not a widget.LayoutView")
	return nil
}
