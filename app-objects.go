package model

type ObjectsHandler interface {
	AddObjects(o ...*Object)
	GetObjects() []*Object
	GetObjectByName(object_name string) (*Object, error)
}

func (h *Handlers) AddObjects(new_objects ...*Object) {

	for _, new := range new_objects {
		for _, o := range h.objects {
			if o.Name == new.Name {
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
		return nil, Error("error nombre objeto no puede estar vació")
	}

	for _, o := range h.objects {
		// d.Log("BUSCANDO OBJETO:", o.Name)
		if o.Name == object_name {
			return o, nil
		}
	}

	return nil, Error("error objeto:", object_name, ",no encontrado")
}
