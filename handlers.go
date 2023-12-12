package model

type Handlers struct {
	ProductionMode    bool
	FrontendExecution bool
	FileRootFolder    string // default "app_files"
	AppInfo

	AuthFrontendAdapter
	AuthBackendAdapter

	BackendBootDataUser
	FrontendBootDataUser

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
}