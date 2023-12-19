package model

func (h *MainHandler) AddModules(new_modules ...*Module) {

	for _, new := range new_modules {
		var exists bool
		for _, m := range h.modules {
			if m.ModuleName == new.ModuleName {
				exists = true
				break
			}
		}

		if !exists {
			h.modules = append(h.modules, new)
		}
	}
}

func (h MainHandler) GetModules() []*Module {
	return h.modules
}
