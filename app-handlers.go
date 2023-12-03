package model

type Handlers struct {
	ProductionMode    bool
	FrontendExecution bool
	FileRootFolder    string // default "app_files"
	EncryptionKey     string // llave de cifrado general
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

	FetchAdapter
	DataConverter

	FileApi
	FileDiskRW

	MessageAdapter
	Logger
}
