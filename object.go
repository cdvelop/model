package model

type Object struct {
	Name           string   //nombre del objeto o tabla
	TextFieldNames []string //nombre de campos mas representativos del objeto o tabla ej: nombre, apellido
	Fields         []Field  //campos del objeto
}
