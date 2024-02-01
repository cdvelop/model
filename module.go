package model

type Module struct {
	//nombre modulo ej: chat,patient,user
	ModuleName string
	//Titulo que vera el usuario ej: "Modulo Fotografía"
	Title string
	// id icono para utilizar en sprite svg ej: icon-info
	IconID string
	//interfaz usuario modulo
	UI
	// ej: keyboard,focus
	FrontendModuleHandlers
	//areas soportadas por el modulo ej: 'a','t','x'
	Areas map[string]string
	// objetos o componentes que contiene el modulo ej: patient,user,datalist,search....
	objects []*Object
	// objeto en uso actualmente en el modulo
	object_actual *Object

	*MainHandler
}

type UI interface {
	UserInterface(u *User) string
}

type ModuleFrontend interface {
}
