package model

const INPUT_PATTERN = `input\.(\w+)\([^)]*\)`

type Input struct {
	InputName string

	Minimum int
	Maximum int

	Tag
	ItemViewAdapter
	ResetViewAdapter
	Validate

	TestData
}

type Tag interface {
	HtmlName() string
	ContainerViewAdapter
}

type Validate interface {
	//options html ej: data-option="ch_doc"
	ValidateField(data_in string, skip_validation bool, options ...string) (err string)
}

type TestData interface {
	GoodTestData() (out []string)
	WrongTestData() (out []string)
}
