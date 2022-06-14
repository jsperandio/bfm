package widget

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	prgCell     = "▉"
	prgMinWidth = 40
)

type ProgressDialog struct {
	*tview.Box
	x       int
	y       int
	width   int
	height  int
	counter int
	display bool
}

func NewProgressDialog() *ProgressDialog {
	return &ProgressDialog{
		Box:     tview.NewBox().SetBorder(true),
		display: false,
		height:  10,
	}
}

// Draw draws this primitive onto the screen.
func (pd *ProgressDialog) Draw(screen tcell.Screen) {
	if !pd.display || pd.height < 3 {
		return
	}
	pd.Box.DrawForSubclass(screen, pd)
	x, y, width, _ := pd.Box.GetInnerRect()
	tickStr := pd.tickStr(width)
	tview.Print(screen, tickStr, x, y, width, tview.AlignLeft, tcell.ColorYellow)

}

// SetRect set rects for this primitive.
func (pd *ProgressDialog) SetRect(x, y, width, height int) {

	pd.x = x
	pd.y = y
	pd.width = width
	if pd.width > prgMinWidth {
		pd.width = prgMinWidth
		spaceWidth := (width - pd.width) / 2
		pd.x = x + spaceWidth
	}
	if height > 3 {
		pd.height = 3
		spaceHeight := (height - pd.height) / 2
		pd.y = y + spaceHeight
	}

	pd.Box.SetRect(pd.x, pd.y, pd.width, pd.height)

}

func (pd *ProgressDialog) Hide() {
	pd.display = false
}

func (pd *ProgressDialog) Display() {
	pd.counter = 0
	pd.display = true
}

func (pd *ProgressDialog) tickStr(max int) string {
	counter := pd.counter
	if counter < max-4 {
		pd.counter++
	} else {
		pd.counter = 0
	}
	prgHeadStr := ""
	hWidth := 0
	prgEndStr := ""
	prgStr := ""
	for i := 0; i < pd.counter; i++ {
		prgHeadStr = prgHeadStr + fmt.Sprintf("[black::]%s", prgCell)
		hWidth++
	}
	prgStr = prgCell + prgCell + prgCell + prgCell
	for i := 0; i < max+hWidth+4; i++ {
		prgEndStr = prgEndStr + fmt.Sprintf("[black::]%s", prgCell)
	}

	progress := fmt.Sprintf("%s[%s::]%s%s", prgHeadStr, "darkorange", prgStr, prgEndStr)
	return progress
}

// InputHandler returns input handler function for this primitive
func (d *ProgressDialog) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return d.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		// log.Printf("progress dialog: event %v received", event)

	})
}
