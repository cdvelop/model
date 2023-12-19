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

	modules []*Module // módulos agregados al sistema
	ObjectHandlerAdapter

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
