package model

type DomAdapter interface {
	BuildFrontendUI() (err string)
	InnerHTML(querySelector, html_content string) (err string)
	InsertAfterBegin(querySelector, html_content string) (err string)
	InsertBeforeEnd(querySelector, html_content string) (err string)
	// querySelector ej: "a[name='xxx']"
	ElementClicking(querySelector string) (err string)
	ElementFocus(querySelector string) (err string)
	ElementsDelete(object_name string, on_server_too bool, objects_ids ...map[string]string) (err string)

	SelectContent(SelectDomOptions struct {
		QuerySelector string // ej: querySelector "meta[name='JsonBootTests']"
		GetContent    string //ej: GetContent: "content"
		SetAfter      bool   // SetAfter ej: true = element.Set("content", "")
		StringReturn  bool   //return_as_string retornar como string de lo contrario sera un jsValue
	}) (jsValue any, err string)

	CallFunction(NameJsFunc string, Options struct {
		Params       map[string]any // Objecto como parámetro
		ResultInt    bool           // resultado en formato int
		ResultString bool           // resultado en string
	}) (result any, err string)
}
