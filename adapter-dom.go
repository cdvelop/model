package model

type DomAdapter interface {
	BuildFrontendUI() (err string)
	InnerHTML(html_content string, o *Object)
	InsertAfterBegin(html_content string, o *Object)
	InsertBeforeEnd(html_content string, o *Object)
	// querySelector ej: "a[name='xxx']"
	ElementClicking(querySelector string) (err string)
	//ej: querySelector "meta[name='JsonBootTests']"
	// get_content: "content"
	SelectContent(querySelector, get_content string) (content, err string)

	CallFunction(functionName string, args ...any) (err string)
}
