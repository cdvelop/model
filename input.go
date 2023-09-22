package model

const INPUT_PATTERN = `input\.(\w+)\([^)]*\)`

type Input struct {
	InputName string

	Minimum int
	Maximum int

	Tag
	Validate

	TestData
}

type Tag interface {
	HtmlName() string
	HtmlTag(id, field_name string, allow_skip_completed bool) string
}

type Validate interface {
	//options html ej: data-option="ch_doc"
	ValidateField(data_in string, skip_validation bool, options ...string) error
}

type TestData interface {
	GoodTestData() (out []string)
	WrongTestData() (out []string)
}
