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

func (m Module) BuildHtmlModule(options ...string) (out string) {
	if m.Theme != nil && m.UI != nil {
		return m.ModuleHtmlTemplate(m.Name, m.UI.UserInterface(options...))
	}
	return
}

func (m Module) BuildMenuButton(index int) (out string) {
	if m.Theme != nil {
		return m.MenuButtonTemplate(m.Name, strconv.Itoa(index), m.IconID, m.Title)
	}
	return
}
