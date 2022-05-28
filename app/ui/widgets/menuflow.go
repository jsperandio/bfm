package widgets

import "github.com/rivo/tview"

type MenuFlow struct {
	*tview.Box
	pages *tview.Pages
	menus map[string]*tview.List
}
