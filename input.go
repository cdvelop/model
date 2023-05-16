package model

type input interface {
	Name() string                                                     //ej radio, text, date, age_date, text_run                                                                  //nombre del input ej: text,radio,date etc...
	Validate(data_in string, skip_validation bool) bool               //como sera validado
	Build(id, field_name string, skip_completion_allowed bool) string //construir vista usuario input
}
