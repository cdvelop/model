package model

type Handlers struct {
	FileRootFolder string // default "app_files"
	AppInfo
	AuthAdapter

	modules []*Module
	objects []*Object

	ThemeAdapter
	DataBaseAdapter
	TimeAdapter
	DomAdapter
	FetchAdapter
	Logger
	DataConverter

	FileApi
	FileDiskRW

	// TEST  TestAdapter
}
