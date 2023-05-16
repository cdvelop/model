package model

type Input struct {
	Name string
	Build
	Validate
	ui
}

type Build interface {
	Build(id, field_name string, skip_completion_allowed bool) string //construir vista usuario input
}

type Validate interface {
	Validate(data_in string, skip_validation bool) bool //como sera validado
}
