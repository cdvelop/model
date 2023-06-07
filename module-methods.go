package model

import "strconv"

func (m Module) ContainsTypeAreas() (public, private bool) {

	for _, a := range m.Areas {

		if a == '0' {
			public = true
		} else {
			private = true
		}

		if public && private {
			break
		}
	}

	return
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
