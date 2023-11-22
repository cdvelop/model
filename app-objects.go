package model

type ObjectsHandler interface {
	AddObjects(o ...*Object)
	GetObjects() []*Object
	GetObjectByName(object_name string) (*Object, error)
}

func (h *Handlers) AddObjects(new_objects ...*Object) {

	for _, new := range new_objects {
		for _, o := range h.objects {
			if o.ObjectName == new.ObjectName {
				return
			}
		}
	}

	h.objects = append(h.objects, new_objects...)

}

func (h Handlers) GetObjects() []*Object {
	return h.objects
}

func (h Handlers) GetObjectByName(object_name string) (*Object, error) {
	// d.Log("total objetos:", len(d.objects))
	if object_name == "" {
		return nil, Error("error GetObjectByName nombre objeto no puede estar vació")
	}

	for _, o := range h.objects {
		// d.Log("BUSCANDO OBJETO:", o.ObjectName)
		if o.ObjectName == object_name {
			return o, nil
		}
	}

	return nil, Error("objeto:", object_name, "no encontrado")
}
