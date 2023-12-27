package model

type Field struct {
	// si el campo comienza com id_ se considera como único y perteneciente a una tabla específica ej: id_user su tabla es user
	// otros ej: id_usuario, apellido, address, city etc
	Name          string
	Legend        string //como se mostrara al usuario el campo en el encabezado ej: "name" por "Nombre Usuario"
	SourceTable   string // tabla origen ej: patient,user,product
	NotRenderHtml bool   // si no se necesita en formulario html

	*Input

	NotClearValueOnFormReset bool
	SkipCompletionAllowed    bool //se permite que el campo no este completado. por defecto siempre sera requerido

	Unique          bool //campo único e inalterable en db
	NotRequiredInDB bool // campo no requerido en base de datos al crear tabla
	Encrypted       bool // si estará encriptado su almacenamiento o no
}

func (f Field) IsPrimaryKey(o *Object) bool {
	if o != nil {
		if o.PrimaryKeyName() == f.Name {
			return true
		}
	}
	return false
}
