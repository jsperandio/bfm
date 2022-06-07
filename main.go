package main

import (
	"github.com/jsperandio/bfm/app/domain/service"
	"github.com/jsperandio/bfm/app/ui"
)

func main() {

	mk := service.NewProjectMaker()

	mainScreen := ui.NewScreen(mk)
	mainScreen.Render()

}
