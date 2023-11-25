package model

type FormAdapter interface {
	FormReset(object_name string) (err string)
	FormAutoFill(object_name string) (err string)
	FormComplete(object_name string, data map[string]string) (err string)
}
