package model

type ModuleHandler interface {
	AddModules(new_modules ...*Module)
	GetModules() []*Module
	GetModuleByName(module_name string) (m *Module, err string)

	SetActualModuleInMainHandler(module_name string) (err string)
	GetActualModuleFromMainHandler() (o *Module, err string)
}

type ObjectsHandler interface {
	GetAllObjectsFromMainHandler() []*Object
	GetObjectByNameMainHandler(object_name string) (o *Object, err string)
}

func (h MainHandler) GetAllObjectsFromMainHandler() (objects []*Object) {
	for _, m := range h.Modules {
		objects = append(objects, m.Objects...)
	}

	return
}

func (h MainHandler) GetObjectByNameMainHandler(object_name string) (o *Object, err string) {
	// d.Log("total objetos:", len(d.objects))

	for _, m := range h.Modules {

		o, err = m.GetObjectFromModule(object_name)
		if o != nil && err == "" { // si no hay error retorno el objeto
			return o, ""
		}
	}

	return nil, object_name + " no encontrado func GetObjectByNameMainHandler"
}

func (h MainHandler) GetModules() []*Module {
	return h.Modules
}

func (h *MainHandler) AddModules(new_modules ...*Module) {

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

func (h *MainHandler) SetActualModuleInMainHandler(module_name string) (err string) {
	h.module_actual, err = h.GetModuleByName(module_name)

	return
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
