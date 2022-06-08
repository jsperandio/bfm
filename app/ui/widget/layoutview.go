package widget

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jsperandio/bfm/app/ui/constant"
	"github.com/jsperandio/bfm/app/ui/model"
	"github.com/rivo/tview"
)

type LayoutView struct {
	*tview.TreeView
	name string
}

func NewLayoutView() *LayoutView {

	lt := &LayoutView{
		TreeView: tview.NewTreeView(),
		name:     constant.LayoutViewName,
	}

	lt.SetBorder(true).SetTitle("Layout Tree")
	lt.SetBorderPadding(1, 0, 2, 0)

	lt.RenderLayout(nil)
	return lt
}

func (lt *LayoutView) RenderLayout(lyt *model.Layout) {

	if lyt == nil {
		return
	}
	node := tview.NewTreeNode(lyt.Name())
	node.SetColor(tcell.ColorGreen)

	lt.SetRoot(node).SetCurrentNode(node)
	lt.nodeDive(node, lyt.Structure["dir"])
}

func (lt *LayoutView) nodeDive(nd *tview.TreeNode, vl interface{}) {

	switch vl := vl.(type) {

	case map[string]interface{}:
		for k, v := range vl {
			n := tview.NewTreeNode(k)
			nd.AddChild(n)
			lt.nodeDive(n, v)
		}
	case string:
		n := tview.NewTreeNode(vl)
		nd.AddChild(n)
	}
}

func (lt *LayoutView) GetName() string {
	return lt.name
}
