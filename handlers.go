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

	modules []*Module //todos los módulos agregados al sistema
	objects []*Object // todos los objetos agregados al sistema
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

	DevicePeripherals

	MessageAdapter
	Logger
}
