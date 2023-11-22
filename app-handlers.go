package model

type Handlers struct {
	ProductionMode bool
	FileRootFolder string // default "app_files"
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

	Test *Tests
}
