package model

type UI interface {
	UserInterface(options ...string) string
}
