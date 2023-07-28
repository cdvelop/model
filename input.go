package model

type Input struct {
	InputName string

	Tag
	Validate

	TestData
}

type Tag interface {
	HtmlName() string
	HtmlTag(id, field_name string, allow_skip_completed bool) string
}

type Validate interface {
	ValidateField(data_in string, skip_validation bool) bool
}

type TestData interface {
	GoodTestData() (out []string)
	WrongTestData() (out []string)
}

const INPUT_PATTERN = `Input: input\.(\w+)\(\)`
