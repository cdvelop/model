package model

type FormAdapter interface {
	FormReset(object_name string) (err string)
	//validate valida después de completar,auto_grow hacer crecer los campos según el contenido
	FormComplete(validate, auto_grow bool) (err string)
}
