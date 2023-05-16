package model

type Component struct {
	//ej: search_footer,datalist,form
	Name string
	//ej: .mi-style{background:black;}
	CSS string
	// ej: MyFunction(e){console.log("hello ",e)};
	FunctionsJS string
	// ej: btn.addEventListener('click', MyFunction);
	ListenerJS string
}
