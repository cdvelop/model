package model

type Module struct {
	Theme
	//nombre modulo ej: chat,patient,user
	Name string
	//Titulo que vera el usuario ej: "Modulo Fotografía"
	Title string
	// id icono para utilizar en sprite svg ej: icon-info
	IconID string
	//interfaz usuario modulo
	UI
	//areas soportadas por el modulo ej: 'a','t','x'
	Areas []byte
	// objetos o componentes que contiene el modulo ej: patient,user,datalist,search....
	Objects []*Object
}

func (m Module) MainName() string {
	return m.Name
}

func (m Module) Components() []*Object {
	return m.Objects
}
