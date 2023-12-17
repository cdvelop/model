package model

type MainHandler struct {
	ProductionMode    bool
	FrontendExecution bool
	FileRootFolder    string // default "app_files"
	AppInfo

	SessionFrontendAdapter
	SessionBackendAdapter

	BackendBootDataUser
	FrontendBootDataUser

	Modules       []*Module
	module_actual *Module

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
