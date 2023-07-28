package model

import "net/http"

type FrontendResponse struct {
}

type BackendRequest struct {
	apiName string //ej staff.calendar,kdks9125=...

	CreateApi
	ReadApi
	UpdateApi
	DeleteApi

	FileApi
}

type CreateApi interface {
	Create(data ...map[string]string) error
}

type ReadApi interface {
	Read(data ...map[string]string) ([]map[string]string, error)
}

type UpdateApi interface {
	Update(data ...map[string]string) ([]map[string]string, error)
}

type DeleteApi interface {
	Delete(data ...map[string]string) ([]map[string]string, error)
}

type FileApi interface {
	MaximumFileSize() int64
	UploadFile(r *http.Request, params map[string]string) ([]map[string]string, error)
	FilePath(params map[string]string) (file_path string, err error)
}
