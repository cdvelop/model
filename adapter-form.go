package model

type FormAdapter interface {
	BuildFormFields() map[string]interface{}
}
