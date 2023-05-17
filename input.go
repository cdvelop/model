package model

type Input struct {
	Component

	Build
	Validate
}

type Build interface {
	HtmlTAG(id, field_name string, skip_completion_allowed bool) string //construir vista usuario input
}

type Validate interface {
	DataField(data_in string, skip_validation bool) bool //como sera validado
}
