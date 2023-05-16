package model

type Component struct {
	//ej: search_footer,datalist,form
	Name string
	//ej: .mi-style{background:black;}
	CssPublic
	CssPrivate

	JsGlobal
	JsModule
	JsListeners
}
