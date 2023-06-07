package model

import "strconv"

type Module struct {
	Theme
	//nombre modulo ej: chat,patient,user
	Name string
	//Titulo que vera el usuario ej: "Modulo Fotografía"
	Title string
	// icono para utilizar en sprite svg
	Icon
	//interfaz usuario modulo
	UI
	//areas soportadas por el modulo ej: 'a','t','x'
	Areas []byte
	// ui/components que usa el modulo ej: form,datalist,search....
	Components []Component
	// objetos que contiene el modulo ej: patient,user
	Objects []Object

	Path
}

func (m Module) BuildIconModule() (out string) {
	if m.Theme != nil {
		return m.IconModuleTemplate(m.Icon.Id, m.Icon.ViewBox, m.Icon.Paths...)
	}
	return
}

func (m Module) BuildHtmlModule(required_area byte) (out string) {
	if m.Theme != nil && m.UI != nil {
		return m.ModuleHtmlTemplate(m.Name, m.UI.UserInterface(required_area))
	}
	return
}

func (m Module) BuildMenuButton(index int) (out string) {
	if m.Theme != nil {
		return m.MenuButtonTemplate(m.Name, strconv.Itoa(index), m.Icon.Id, m.Title)
	}
	return
}

func (m Module) BuildModuleJS(functions, listener_add, listener_rem string) (out string) {
	if m.Theme != nil {
		return m.ModuleJsTemplate(m.Name, functions, listener_add, listener_rem)
	}
	return
}
