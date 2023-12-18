package model

type ObjectHandler interface {
	ObjectActual() *Object
	SetActualObject(object_name string) (err string)
	MainHandlerGetAllObjects() (objects []*Object)
	MainHandlerGetObjectByName(object_name string) (o *Object, err string)
}

type ModuleHandler interface {
	// MainHandlerAddModules(new_modules ...*Module)
	// MainHandlerGetModules() []*Module
	// GetModuleByName(module_name string) (m *Module, err string)

	// GetActualModuleFromMainHandler() (o *Module, err string)
}

func (h MainHandler) ObjectActual() *Object {
	return h.module_actual.object_actual
}

// MainHandlerGetAllObjects() []*Object
