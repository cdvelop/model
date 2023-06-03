package model

type Input struct {
	Component

	Tag
	Validate

	TestData
}

type Tag interface {
	HtmlTag(id, field_name string, allow_skip_completed bool) string
}

type Validate interface {
	ValidateField(data_in string, skip_validation bool) bool
}

type TestData interface {
	GoodTestData() (out []string)
	WrongTestData() (out []string)
}
