package model

type Logger interface {
	// pare usar Log en una lines en funciones de tipo syscall/js se usa interface{}
	Log(message ...any)
}
