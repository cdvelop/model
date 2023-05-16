package model

type Component struct {
	//ej: search_footer,datalist,form
	Name string
	//ej: .mi-style{background:black;}
	CSS string
	JS
}

type JS struct {
	// ej: MyFunction(e){console.log("hello ",e)};
	Functions string
	// ej: btn.addEventListener('click', MyFunction);
	Listeners string
}
