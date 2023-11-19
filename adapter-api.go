package model

type BackendHandler struct {
	BootResponse

	CreateApi
	ReadApi
	UpdateApi
	DeleteApi
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

type ReadData interface {
	ReadByID(id string) ([]map[string]string, error)
}

type UpdateApi interface {
	Update(u *User, data ...map[string]string) error
}

type DeleteApi interface {
	Delete(u *User, data ...map[string]string) ([]map[string]string, error)
}
