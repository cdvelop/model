package model

type ModuleHandler interface {
	GetModules() []*Module
	GetModuleByName(module_name string) (m *Module, err string)
}

func (h Handlers) GetModules() []*Module {
	return h.modules
}

func (h *Handlers) AddModules(new_modules ...*Module) {

	for _, new := range new_modules {
		for _, m := range h.modules {
			if m.ModuleName == new.ModuleName {
				return
			}
		}
	}

	h.modules = append(h.modules, new_modules...)
}

func (h Handlers) GetModuleByName(module_name string) (m *Module, err string) {
	const this = "GetModuleByName error "
	// d.Log("total objetos:", len(d.objects))
	if module_name == "" {
		return nil, this + "nombre modulo no puede estar vació"
	}

	for _, m := range h.modules {
		// d.Log("BUSCANDO OBJETO:", o.ObjectName)
		if m.ModuleName == module_name {
			return m, ""
		}
	}

	return nil, this + "modulo: " + module_name + ", no encontrado"
}
