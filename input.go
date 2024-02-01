package model

const INPUT_PATTERN = `input\.(\w+)\([^)]*\)`

type Input struct {
	InputName string

	Tag
	ItemViewAdapter
	Validate

	ResetParameters *struct {
		ResetJsFuncName    string         // nombre la la función js a llamar para reset
		Enable             bool           //parámetro estado por defecto a enviar
		NotSendQueryObject bool           //no envía querySelector Objeto por defecto siempre se envía
		Params             map[string]any // parámetros adicionales
	}

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
