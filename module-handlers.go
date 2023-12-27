package model

type ClickedModuleEventAdapter interface {
	ClickedModuleEvent()
}

type FrontendModuleHandlers struct {
	ClickedModuleEventAdapter
}
