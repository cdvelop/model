package model

type Component struct {
	//ej: search_footer,datalist,form
	Name string
	//Css Global ej: .mi-style{background:black;}
	CssGlobal
	//Css Privado
	CssPrivate
	// Js Global
	JsGlobal
	// Js Privado
	JsPrivate
	//Js Listeners
	JsListeners
}
