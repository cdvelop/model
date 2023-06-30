package model

type FrontEnd struct {
}

type ApiHandler struct {
	// nombre del componente u objeto ej: client, search_footer,datalist,form
	Name string

	CreateApi

	ReadFileApi
	ReadOneApi
	ReadAllApi

	UpdateApi
	DeleteApi
}

type CreateApi interface {
	Create(params *[]map[string]string) error
}

type ReadFileApi interface {
	ReadFile(params *map[string]string) (file_path string, err error)
}

type ReadOneApi interface {
	ReadOne(params *map[string]string) (*map[string]string, error)
}

type ReadAllApi interface {
	ReadAll(params *map[string]string) (*[]map[string]string, error)
}

type UpdateApi interface {
	Update(params *[]map[string]string) error
}

type DeleteApi interface {
	Delete(params *[]map[string]string) error
}
