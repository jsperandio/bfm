package constant

import "time"

const (
	// General
	ScreenRefreshRate = time.Millisecond * 50
	FileViewName      = "FileView"
	LayoutViewName    = "LayoutView"
	ModalProgress     = "ModalProgress"

	ScreenMenuProportion = 3
	ScreenViewProportion = 3

	// Main Menu Constants
	MainMenuTitle = "BFM - Build for Me - v1.0"
	MainMenuName  = "MainMenu"

	// Project Menu Constants
	ProjectMenuTitle = "New Project"
	ProjectMenuName  = "NewProjectMenu"
	KeyboardRunes    = "qwertyuiopasdfghjklzxcvnm" // removed [b] for back option

	// Param Form Constants
	ParamFormName                  = "ParamForm"
	ParamFormInitialPathDefaultDir = "./documents/repository"
	ParamFormGitUserNameDefault    = "johndoe"
)
