package model

type Object struct {
	Name           string   //nombre del objeto o tabla
	TextFieldNames []string //nombre de campos mas representativos del objeto o tabla ej: nombre, apellido
	Fields         []Field  //campos del objeto
	//operaciones permitidas según nivel de acceso ej: 0,1,2,3,4 + Create bool, Update bool, Delete bool
	OperationsAllowed map[uint8]Permission
}
