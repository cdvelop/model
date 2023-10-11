package model

type DomAdapter interface {
	Log
	UserMessage

	InsertAfterBegin(...ViewHandler)
	InsertBeforeEnd(...ViewHandler)

	Clicking(o *Object) error
}

type Log interface {
	// retorno una interfaz solo para simplificar funciones de tipo syscall/js
	Log(message ...any) interface{}
}

type UserMessage interface {
	// retorno una interfaz solo para simplificar funciones de tipo syscall/js
	UserMessage(message ...any) interface{}
}
