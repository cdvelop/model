package model

func (m *Module) AddObjectsToModule(new_objects ...*Object) {

	for _, new := range new_objects {
		exists := false
		for _, o := range m.objects {
			if o.ObjectName == new.ObjectName {
				exists = true
				break
			}
		}

		if !exists {
			m.objects = append(m.objects, new)
		}
	}
}
func (m Module) GetObjects() []*Object {
	return m.objects
}

// obtener todos los objetos del modulo
func (m Module) GetObject(object_name string) (o *Object, err string) {
	const e = " . GetObject"

	if object_name == "" {
		return nil, "nombre no puede estar vació" + e
	}

	for _, o := range m.objects {
		// d.Log("BUSCANDO OBJETO:", o.ObjectName)
		if o.ObjectName == object_name {
			return o, ""
		}
	}

	return nil, "objeto:" + object_name + ", no pertenece al modulo:" + m.ModuleName + e
}

// obtener objeto en uso actualmente en el modulo
func (m *Module) GetActualModuleObject() (o *Object, err string) {

	if m.object_actual == nil {
		return nil, "objecto actual no definido func GetActualModuleObject"
	}

	return m.object_actual, ""
}

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

	for _, o := range m.objects {
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
