package model

type JsGlobal interface {
	JsG() string
}

type JsPrivate interface {
	JsP() string
}

type JsListeners interface {
	JsL() string
}
