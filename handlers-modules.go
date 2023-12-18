package model

func (h MainHandler) MainHandlerGetAllObjects() (objects []*Object) {
	for _, m := range h.Modules {
		objects = append(objects, m.Objects...)
	}

	return
}

func (h MainHandler) MainHandlerGetObjectByName(object_name string) (o *Object, err string) {
	// d.Log("total objetos:", len(d.objects))

	for _, m := range h.Modules {

		o, err = m.GetObjectFromModule(object_name)
		if o != nil && err == "" { // si no hay error retorno el objeto
			return o, ""
		}
	}

	return nil, object_name + " no encontrado func MainHandlerGetObjectByName"
}

func (h MainHandler) MainHandlerGetModules() []*Module {
	return h.Modules
}

func (h *MainHandler) MainHandlerAddModules(new_modules ...*Module) {

	for _, new := range new_modules {
		exists := false
		for _, m := range h.Modules {
			if m.ModuleName == new.ModuleName {
				exists = true
				break
			}
		}

		if !exists {
			h.Modules = append(h.Modules, new)
		}
	}
}

func (h *MainHandler) SetActualModule(module_name string) (err string) {

	if h.module_actual != nil && h.module_actual.ModuleName == module_name {
		return
	}

	h.module_actual, err = h.GetModuleByName(module_name)

	return
}

func (h *MainHandler) SetActualObject(object_name string) (err string) {
	return h.module_actual.setActualModuleObject(object_name)
}

func (h MainHandler) GetModuleByName(module_name string) (m *Module, err string) {
	const e = "func GetModuleByName"
	// d.Log("total objetos:", len(d.objects))
	if module_name == "" {
		return nil, "nombre modulo no puede estar vació" + e
	}

	m, err = h.GetActualModuleFromMainHandler()
	if m != nil && err == "" && m.ModuleName == module_name {
		return m, ""
	}

	for _, m := range h.Modules {
		// d.Log("BUSCANDO OBJETO:", o.ObjectName)
		if m.ModuleName == module_name {
			return m, ""
		}
	}

	return nil, "modulo: " + module_name + ", no encontrado" + e
}

// obtener modulo en uso actualmente por el usuario
func (h MainHandler) GetActualModuleFromMainHandler() (o *Module, err string) {
	if h.module_actual == nil {
		return nil, "modulo actual no definido func GetActualModuleFromMainHandler"
	}
	return h.module_actual, ""
}
