package model

type Component struct {
	//ej: search_footer,datalist,form
	Name string
	// CssGlobal() string | Css Global ej: .mi-style{background:black;}
	CssGlobal
	// CssPrivate() string | Css Privado
	CssPrivate
	// JsGlobal() string | Js Global
	JsGlobal
	// JsPrivate() string | Js Privado Modulo ej: MyFunction(e){console.log("hello ",e)};
	JsPrivate
	// JsListeners() string | ej: btn.addEventListener('click', MyFunction);
	JsListeners
}
