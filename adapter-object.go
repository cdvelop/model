package model

type ObjectHandlerAdapter interface {
	ObjectActual() *Object
	GetAllObjects() (objects []*Object)
}
