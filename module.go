package model

type Module struct {
	//nombre modulo ej: chat,patient,user
	Name string
	//Titulo que vera el usuario ej: "Modulo Fotografía"
	Title string
	// icono según sprite svg ej: "icon-camera"
	Icon string

	UI
	//areas soportadas por el modulo ej: 'a','t','x'
	Areas []byte
	// ui/components que usa el modulo ej: form,datalist,search....
	Components []Component
	// objetos que contiene el modulo ej: patient,user
	Objects []Object
}
