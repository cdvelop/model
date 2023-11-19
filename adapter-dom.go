package model

type DomAdapter interface {
	InnerHTML(html_content string, o *Object)
	InsertAfterBegin(html_content string, o *Object)
	InsertBeforeEnd(html_content string, o *Object)

	ClickModule(module_name string) error
	Clicking(o *Object, id string) error

	CallFunction(functionName string, args ...any) error
}
