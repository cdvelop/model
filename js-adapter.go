package model

type JS interface {
	//Js Global
	JsGlobal() string
	//Js Modulo ej: MyFunction(e){console.log("hello ",e)};
	JsModule() string
	// ej: btn.addEventListener('click', MyFunction);
	JsListeners() string
}
