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

func (m Module) BuildSpriteIcon() (out string) {
	if m.Theme != nil {
		return m.SpriteIconTemplate(m.Icon.Id, m.Icon.ViewBox, m.Icon.Paths...)
	}
	return
}

func (m Module) BuildHtmlModule(options ...string) (out string) {
	if m.Theme != nil && m.UI != nil {
		return m.ModuleHtmlTemplate(m.MainName, m.UI.UserInterface(options...))
	}
	return
}

func (m Module) BuildMenuButton(index int) (out string) {
	if m.Theme != nil {
		return m.MenuButtonTemplate(m.MainName, strconv.Itoa(index), m.Icon.Id, m.Title)
	}
	return
}

func (m Module) BuildModuleJS(functions, listener_add, listener_rem string) (out string) {
	if m.Theme != nil {
		return m.ModuleJsTemplate(m.MainName, functions, listener_add, listener_rem)
	}
	return
}

func (m *Module) AddObjects(objects ...*Object) {
	for _, new_obj := range objects {
		if new_obj != nil {

			var obj_found bool

			for _, o := range m.Objects {
				if o.Name == new_obj.Name {
					obj_found = true
					break
				}
			}

			if !obj_found {
				m.Objects = append(m.Objects, new_obj)
				new_obj.addModule(m)
			}

		}
	}
}

func (o *Object) addModule(new_module *Module) {
	if new_module != nil {

		var module_found bool

		for _, m := range o.Modules {
			if m.MainName == new_module.MainName {
				module_found = true
				break
			}
		}

		if !module_found {
			o.Modules = append(o.Modules, new_module)
		}

	}
}
