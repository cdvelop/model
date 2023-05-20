package model

type Input struct {
	Component

	Build
	Validate

	TestData
}

type Build interface {
	HtmlTAG(id, field_name string, allow_skip_completed bool) string //construir vista usuario input
}

type Validate interface {
	ValidateField(data_in string, skip_validation bool) bool //como sera validado
}

type TestData interface {
	GoodTestData(table_name, field_name string, random bool) (out []string)
	WrongTestData() (out []string)
}
