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

	modules    []*Module //todos los módulos agregados al sistema
	objects    []*Object // todos los objetos agregados al sistema
	objects_db []*Object // todos los objetos de tipo tabla en base de datos

	ObjectHandlerAdapter

	ThemeAdapter
	DataBaseAdapter
	BackupHandlerAdapter
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
