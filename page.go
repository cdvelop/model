package model

// index.html ej:
// {{.StyleSheet}} {{.AppName}} {{.AppVersion}} {{.SpriteIcons}} {{.Menu}}
// {{.Message}} {{.UserName}} {{.UserArea}} {{.Modules}} {{.Script}}
type Page struct {
	StyleSheet string // url ej style.css

	AppName    string
	AppVersion string

	SpriteIcons string

	Menu string

	UserName string
	UserArea string
	Message  string

	Modules string

	Script string // ej main.js
}
