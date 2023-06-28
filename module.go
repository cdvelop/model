package model

type Module struct {
	Theme
	//nombre modulo ej: chat,patient,user
	MainName string
	//Titulo que vera el usuario ej: "Modulo Fotografía"
	Title string
	// icono para utilizar en sprite svg
	Icon
	//interfaz usuario modulo
	UI
	//areas soportadas por el modulo ej: 'a','t','x'
	Areas []byte
	// objetos o componentes que contiene el modulo ej: patient,user,datalist,search....
	Objects []*Object

	Path
}
