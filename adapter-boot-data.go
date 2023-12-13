package model

type BackendBootDataUser interface {
	BackendLoadBootData(u *User) (out, err string)
}

type FrontendBootDataUser interface {
	FrontendLoadHtmlBootData(data string) (err string)
}
