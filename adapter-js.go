package model

type JsGlobal interface {
	//Js Global
	JsGlobal() string
}

type JsModule interface {
	//Js Modulo ej: MyFunction(e){console.log("hello ",e)};
	JsModule() string
}

type JsListeners interface {
	// ej: btn.addEventListener('click', MyFunction);
	JsListeners() string
}
