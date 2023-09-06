package model

import "net/http"

type FrontendResponse struct {
}

type BackendRequest struct {
	BootResponse

	CreateApi
	ReadApi
	UpdateApi
	DeleteApi

	FileApi
}

type BootResponse interface {
	AddBootResponse(u *User) ([]Response, error)
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
	CreateFile(r *http.Request, params map[string]string) ([]map[string]string, error)
	FilePath(params map[string]string) (file_path string, err error)
}
