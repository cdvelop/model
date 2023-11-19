package model

type FormAdapter interface {
	FormReset(o *Object) error
	FormAutoFill(o *Object) error
	FormComplete(o *Object, data map[string]string) error
}
