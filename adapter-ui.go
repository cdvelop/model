package model

type ui interface {
	CssPublic
	CssPrivate
	JsGlobal
	JsModule
	JsListeners
}
