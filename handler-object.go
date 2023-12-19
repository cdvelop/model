package model

type ObjectHandlerAdapter interface {
	ObjectActual() *Object
	SetActualObject(object_name string) (err string)
}

type ObjectsHandlerAdapter interface {
	GetObjectByName(name string) (obj *Object, err string)
	GetAllObjects() (objects []*Object)
}

func (h *MainHandler) GetAllObjects() []*Object {

	if len(h.objects) != 0 {
		return h.objects
	}

	for _, m := range h.GetModules() {
		h.objects = append(h.objects, m.GetObjects()...)
	}

	return h.objects

}
func (h *MainHandler) GetObjectByName(name string) (obj *Object, err string) {

	for _, o := range h.GetAllObjects() {
		if o.ObjectName == name {
			return o, ""
		}
	}

	return nil, "objeto:" + name + " no encontrado."

}
