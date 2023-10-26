package model

type Logger interface {
	// retorno una interfaz solo para simplificar funciones de tipo syscall/js
	Log(message ...any) interface{}
}
