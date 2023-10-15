package model

type DomAdapter interface {
	Log
	UserMessage

	InnerHTML(html_content string, o *Object)
	InsertAfterBegin(html_content string, o *Object)
	InsertBeforeEnd(html_content string, o *Object)

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
