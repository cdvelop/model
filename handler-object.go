package model

type ObjectHandlerAdapter interface {
	ObjectActual() *Object
	SetActualObject(object_name string) (err string)
}

type ObjectsHandlerAdapter interface {
	// optional object_name or table_name ej: file.user,"" or "","user"
	GetObjectBY(object_name, table_name string) (o *Object, err string)
	GetAllObjects(only_type_db_table bool) (objects []*Object)
}

// only_type_db_table de tipo tabla en base de datos ej: user,patient
func (h *MainHandler) GetAllObjects(only_type_db_table bool) []*Object {
	h.addNormalObjects()
	h.addDbTableObjects()

	if only_type_db_table {
		return h.objects_db
	}

	return h.objects
}

func (h *MainHandler) addNormalObjects() {

	if len(h.objects) == 0 {
		for _, m := range h.GetModules() {
			h.objects = append(h.objects, m.GetObjects()...)
		}
	}
}

func (h *MainHandler) addDbTableObjects() {
	if len(h.objects_db) == 0 {
		for _, o := range h.objects {
			if !o.NoAddObjectInDB && len(o.Fields) != 0 {
				h.objects_db = append(h.objects_db, o)
			}
		}
	}
}

// optional object_name or table_name ej: file.user,"" or "","user"
func (h MainHandler) GetObjectBY(object_name, table_name string) (obj *Object, err string) {

	for _, o := range h.GetAllObjects(false) {
		if object_name != "" && o.ObjectName == object_name || table_name != "" && o.Table == table_name {
			return o, ""
		}
	}

	return nil, "GetObjectBY: " + object_name + table_name + " no encontrado."

}
