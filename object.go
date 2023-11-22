package model

type Object struct {
	// ej: module.client, client.search_footer,user.datalist
	ObjectName string
	Table      string //tabla origen db ej: users

	PrincipalFieldsName []string //campos mas representativos ej: name, address, phone
	NoAddObjectInDB     bool
	Fields              []Field //campos

	*Module // módulo origen

	BackHandler BackendHandler

	FrontHandler FrontendHandler

	PrinterHandler

	FormData map[string]string // data temporal formulario

	AlternativeValidateAdapter
}

func (o Object) MainName() string {
	return o.ObjectName
}

// options ej: update,delete,"mensaje a enviar" default action: create
func (o Object) Response(data []map[string]string, options ...string) Response {

	var execute = "create"
	var message string
	for _, opt := range options {
		switch opt {
		case "update", "delete", "error", "test":
			execute = opt
		default:
			message = opt
		}
	}

	return Response{
		Action:  execute,
		Object:  o.ObjectName,
		Message: message,
		Data:    data,
	}
}
