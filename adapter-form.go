package model

type FormAdapter interface {
	FormReset(o *Object) (err string)
	FormAutoFill(o *Object) (err string)
	FormComplete(o *Object, data map[string]string) (err string)
}
