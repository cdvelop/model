package model

import "net/http"

type BackendHandler struct {
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
	Create(u *User, data ...map[string]string) error
}

type ReadApi interface {
	Read(u *User, data ...map[string]string) ([]map[string]string, error)
}

type UpdateApi interface {
	Update(u *User, data ...map[string]string) ([]map[string]string, error)
}

type DeleteApi interface {
	Delete(u *User, data ...map[string]string) ([]map[string]string, error)
}

type FileApi interface {
	MaximumFileSize() int64
	CreateFile(u *User, r *http.Request, params map[string]string) ([]map[string]string, error)
	FilePath(u *User, params map[string]string) (file_path string, err error)
}
