package model

type MessageAdapter interface {
	UserMessage(message ...any)
}
