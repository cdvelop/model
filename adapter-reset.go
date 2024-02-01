package model

type ResetFrontendObjectStateAdapter interface {
	ResetFrontendObjectState() (err string)
}

type ResetViewHandlerObjectAdapter interface {
	ResetViewHandlerObject() (err string)
}
