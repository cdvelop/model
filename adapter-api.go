package model

type FrontEnd struct {
}

type ApiHandler struct {
	// nombre del componente u objeto ej: client, search_footer,datalist,form
	Name string

	CreateApi
	ReadApi
	UpdateApi
	DeleteApi
}

type CreateApi interface {
	Create(data *[]map[string]string) error
}

type ReadApi interface {
	Read(data *[]map[string]string) error
}

type UpdateApi interface {
	Update(data *[]map[string]string) error
}

type DeleteApi interface {
	Delete(data *[]map[string]string) error
}
