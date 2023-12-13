package model

type DomAdapter interface {
	BuildFrontendUI() (err string)
	InnerHTML(querySelector, html_content string) (err string)
	InsertAfterBegin(querySelector, html_content string) (err string)
	InsertBeforeEnd(querySelector, html_content string) (err string)
	// querySelector ej: "a[name='xxx']"
	ElementClicking(querySelector string) (err string)
	//ej: querySelector "meta[name='JsonBootTests']"
	// get_content: "content"
	// set_after true = element.Set("content", "")
	SelectContent(querySelector, get_content string, set_after bool) (content, err string)

	CallFunction(functionName string, args ...any) (err string)
}
