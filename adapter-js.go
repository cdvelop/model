package model

type JsGlobal interface {
	JsGlobal() string
}

type JsPrivate interface {
	JsPrivate() string
}

type JsListeners interface {
	JsListeners() string
}
