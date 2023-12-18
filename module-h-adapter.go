package model

type ModuleHandlerAdapter interface {
	GetActualModuleObject() (o *Object, err string)
}
