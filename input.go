package model

const INPUT_PATTERN = `input\.(\w+)\([^)]*\)`

type Input struct {
	InputName string

	Minimum int
	Maximum int

	Tag
	InputView
	InputReset
	Validate

	TestData
}

type Tag interface {
	HtmlName() string
	HtmlTag(id, field_name string, allow_skip_completed bool) string
}

type InputReset interface {
	ResetInput()
}

type InputView interface {
	BuildNewView(new_data []map[string]string) (html string)
}

type Validate interface {
	//options html ej: data-option="ch_doc"
	ValidateField(data_in string, skip_validation bool, options ...string) error
}

type TestData interface {
	GoodTestData() (out []string)
	WrongTestData() (out []string)
}
