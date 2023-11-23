package model

type BackendHandler struct {
	BootResponse

	CreateApi
	ReadApi
	UpdateApi
	DeleteApi
}

type BootResponse interface {
	AddBootResponse(u *User) (r []Response, err string)
}

type CreateApi interface {
	Create(u *User, data ...map[string]string) (err string)
}

type ReadApi interface {
	Read(u *User, data ...map[string]string) (out []map[string]string, err string)
}

type ReadData interface {
	ReadByID(id string) (out []map[string]string, err string)
}

type UpdateApi interface {
	Update(u *User, data ...map[string]string) (err string)
}

type DeleteApi interface {
	Delete(u *User, data ...map[string]string) (err string)
}
