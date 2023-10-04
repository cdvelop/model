package model

type DomAdapter interface {
	Log(message ...any)
	UserMessage(text string, options ...string)

	InsertAfterBegin(...ViewHandler)
	InsertBeforeEnd(...ViewHandler)
}
