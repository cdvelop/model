package model

type FormAdapter interface {
	// limpia los valores del formulario y campos en el map del objeto
	// nombre de objeto opcional de lo contrario se usara el actual
	FormClean(object_name ...string) (err string)
	//validate valida después de completar,auto_grow hacer crecer los campos según el contenido
	FormComplete(object_name string, data map[string]string, validate, auto_grow bool) (err string)
	// si no se envía el nombre el campo se realiza foco en el primer elemento
	FormInputFocus(object_name string, focus_field_name ...string) (err string)
}
