package model

var (
	app_version    = ""
	go_version     = ""
	tinygo_version = ""
)

type AppInfo struct {
	App_name         string
	Business_name    string
	Business_address string
	Business_phone   string
}

func (AppInfo) AppVersion() string {
	return app_version
}
func (AppInfo) GoVersion() string {
	return go_version
}

func (AppInfo) TinyGoVersion() string {
	return tinygo_version
}
