package model

type HtmlAdapter interface {
	// retorna un elemento html del tipo js.Value
	GetHtmlModule(module_name string) (jsValue any, err string)
}
