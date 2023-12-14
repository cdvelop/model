package model

type DomAdapter interface {
	BuildFrontendUI() (err string)
	InnerHTML(querySelector, html_content string) (err string)
	InsertAfterBegin(querySelector, html_content string) (err string)
	InsertBeforeEnd(querySelector, html_content string) (err string)
	// querySelector ej: "a[name='xxx']"
	ElementClicking(querySelector string) (err string)

	SelectContent(SelectDomOptions) (jsValue any, err string)

	CallFunction(NameJsFunc string, params ...any) (result any, err string)
}

type SelectDomOptions struct {
	QuerySelector string // ej: querySelector "meta[name='JsonBootTests']"
	GetContent    string //ej: GetContent: "content"
	SetAfter      bool   // SetAfter ej: true = element.Set("content", "")
	StringReturn  bool   //return_as_string retornar como string de lo contrario sera un jsValue
}
