package model

type Field struct {
	// si el campo comienza com id_ se considera como único y perteneciente a una tabla específica ej: id_user su tabla es user
	// otros ej: id_usuario, apellido, address, city etc
	Name                  string
	Unique                bool   //campo único en db
	Inalterable           bool   //la data en el campo sera inalterable después de la creación (¡no se puede modificar!)
	Legend                string //como se mostrara al usuario el campo en el encabezado ej: "name" por "Nombre Usuario"
	Input                 input
	SkipValidation        bool //sin validar
	SkipCompletionAllowed bool //si el campo es requerido obligatoriamente o no

	NotRenderHtml bool // si no se necesita en formulario html

}
