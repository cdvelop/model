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
