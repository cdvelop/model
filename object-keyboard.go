package model

type KeyboardHandlerObject struct {
	KeyEnterAdapter
}

type KeyEnterAdapter interface {
	KeyEnter()
}
