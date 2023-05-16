package model

type JsGlobal interface {
	//Js Global
	JsG() string
}

type JsPrivate interface {
	//Js Modulo ej: MyFunction(e){console.log("hello ",e)};
	JsP() string
}

type JsListeners interface {
	// ej: btn.addEventListener('click', MyFunction);
	JsL() string
}
