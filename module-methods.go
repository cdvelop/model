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
