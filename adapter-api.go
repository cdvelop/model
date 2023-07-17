package model

type FrontendResponse struct {
}

type BackendRequest struct {
	apiName string //ej staff.calendar,kdks9125=...
	CreateApi

	ReadOneApi
	ReadAllApi

	FileApi

	UpdateApi
	DeleteApi
}

type CreateApi interface {
	Create(params *[]map[string]string) (*[]map[string]string, error)
}

type ReadOneApi interface {
	ReadOne(params *map[string]string) (*map[string]string, error)
}

type ReadAllApi interface {
	ReadAll(params *map[string]string) (*[]map[string]string, error)
}

type FileApi interface {
	FilePath(params *map[string]string) (file_path string, err error)
}

type UpdateApi interface {
	Update(params *[]map[string]string) error
}

type DeleteApi interface {
	Delete(params *[]map[string]string) error
}
