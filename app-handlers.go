package model

type Handlers struct {
	ProductionMode    bool
	FrontendExecution bool
	FileRootFolder    string // default "app_files"
	AppInfo
	AuthAdapter

	modules []*Module
	objects []*Object

	ThemeAdapter
	DataBaseAdapter
	TimeAdapter

	DomAdapter
	FormAdapter
	HtmlAdapter

	CipherAdapter
	DataConverter
	FetchAdapter

	FileApi
	FileDiskRW

	MessageAdapter
	Logger

	BackendBootDataUser
	FrontendBootDataUser
}
