package model

type ModuleHandler interface {
	GetModules() []*Module
	GetModuleByName(module_name string) (*Module, error)
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

func (h Handlers) GetModuleByName(module_name string) (*Module, error) {
	// d.Log("total objetos:", len(d.objects))
	if module_name == "" {
		return nil, Error("error nombre modulo no puede estar vació")
	}

	for _, m := range h.modules {
		// d.Log("BUSCANDO OBJETO:", o.Name)
		if m.ModuleName == module_name {
			return m, nil
		}
	}

	return nil, Error("modulo:", module_name, ", no encontrado")
}
