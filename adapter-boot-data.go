package model

type BootPageData struct {
	JsonBootActions string
}

type BackendBootDataUser interface {
	BackendLoadBootData(u *User) BootPageData
}

type FrontendBootDataUser interface {
	FrontendLoadHtmlBootData(data string) (err string)
}
