package model

type Component struct {
	//ej: search_footer,datalist,form
	Name string
	// CssG() string | Css Global ej: .mi-style{background:black;}
	CssGlobal
	// CssP() string | Css Privado
	CssPrivate
	// JsG() string | Js Global
	JsGlobal
	// JsP() string | Js Privado Modulo ej: MyFunction(e){console.log("hello ",e)};
	JsPrivate
	// JsL() string | ej: btn.addEventListener('click', MyFunction);
	JsListeners
}
