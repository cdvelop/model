package model

type JS interface {
	// ej: MyFunction(e){console.log("hello ",e)};
	Functions() string
	// ej: btn.addEventListener('click', MyFunction);
	Listeners() string
}
