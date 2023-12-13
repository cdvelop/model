package model

// from ej: file, user
func (m *Module) AddInputs(new_inputs []*Input, from string) (err string) {
	for _, new := range new_inputs {

		if new == nil {
			return "AddInputs error input nulo al agregar en modulo:" + m.ModuleName + "desde:" + from
		}

		var found bool
		for _, old := range m.Inputs {
			if new.InputName == old.InputName {
				found = true
			}
		}

		if !found {
			m.Inputs = append(m.Inputs, new)
		}

	}

	return ""
}

func (m Module) ModuleSupports(area string) bool {

	if len(m.Areas) != 0 { // hay areas definidas
		// si no lo soporta continuamos a la siguiente
		if _, supports := m.Areas[area]; !supports {
			return false
		}

	}

	return true
}

func (m Module) ResetFrontendStateObjects() {

	for _, o := range m.Objects {
		// m.Log("info reset Object:", o.ObjectName)
		if o.FrontHandler.ResetFrontendObjectStateAdapter != nil {
			// m.Log("ok ResetFrontendObjectState:", o.ObjectName)
			m.Log(o.FrontHandler.ResetFrontendObjectState())
		}

		if o.FrontHandler.ObjectViewHandler != nil {
			// m.Log("ResetViewHandlerObject:", o.ObjectName)
			m.Log(o.FrontHandler.ObjectViewHandler.ResetViewHandlerObject())
		}
	}
}
