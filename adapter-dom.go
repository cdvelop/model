package model

type DomAdapter interface {
	UserMessage

	InnerHTML(html_content string, o *Object)
	InsertAfterBegin(html_content string, o *Object)
	InsertBeforeEnd(html_content string, o *Object)

	Clicking(o *Object, id string) error

	ClickModule(module_name string) error
	CallFunction(functionName string, args ...any) error

	Form
}

type Form interface {
	FormReset(o *Object) error
	FormAutoFill(o *Object) error
	FormComplete(o *Object, data map[string]string) error
}

type UserMessage interface {
	// retorno una interfaz solo para simplificar funciones de tipo syscall/js
	UserMessage(message ...any) interface{}
}
