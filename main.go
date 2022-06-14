package main

import (
	"github.com/jsperandio/bfm/app/domain/service"
	"github.com/jsperandio/bfm/app/ui"
)

func main() {

	gm := service.NewGoInitializer()
	mk := service.NewProjectMaker(gm)

	mainScreen := ui.NewScreen(mk)
	mainScreen.Render()
}
