package model

type DomAdapter interface {
	InnerHTML(html_content string, o *Object)
	InsertAfterBegin(html_content string, o *Object)
	InsertBeforeEnd(html_content string, o *Object)
	// querySelector ej: "a[name='xxx']"
	ElementClicking(querySelector string) (err string)

	CallFunction(functionName string, args ...any) (err string)
	//esperar en milliseconds ej: 500, 2000 ..
	WaitFor(milliseconds int, callback func())
}
