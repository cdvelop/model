package model

type ResetFrontendObjectStateAdapter interface {
	ResetFrontendObjectState() (err string)
}

type ResetViewHandlerObjectAdapter interface {
	ResetViewHandlerObject() (err string)
}

type ResetInputViewAdapter interface {
	ResetInputView() (err string)
}
