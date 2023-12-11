package model

type ObjectsHandler interface {
	AddObjects(o ...*Object)
	GetObjects() []*Object
	GetObjectByName(object_name string) (o *Object, err string)
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

func (h Handlers) GetObjectByName(object_name string) (o *Object, err string) {
	const t = "GetObjectByName objeto: "
	// d.Log("total objetos:", len(d.objects))
	if object_name == "" {
		return nil, t + "nombre no puede estar vació"
	}

	for _, o := range h.objects {
		// d.Log("BUSCANDO OBJETO:", o.ObjectName)
		if o.ObjectName == object_name {
			return o, ""
		}
	}

	return nil, t + object_name + " no encontrado"
}
