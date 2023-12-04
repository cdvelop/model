package model

type BackendBootDataUser interface {
	BackendLoadBootData(u *User) (out, err string)
}

type FrontendBootDataUser interface {
	FrontendLoadBootData(data string) (err string)
}
