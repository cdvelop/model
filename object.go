package model

type Object struct {
	// ej: module.client, client.search_footer,user.datalist
	Name  string
	Table string //tabla origen db ej: users

	PrincipalFieldsName []string //campos mas representativos ej: name, address, phone
	Fields              []Field  //campos

	*Module // módulo origen

	BackendHandler

	FrontendHandler

	PrinterHandler

	FormData map[string]string // data temporal formulario
}

func (o Object) MainName() string {
	return o.Name
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
		Object:  o.Name,
		Message: message,
		Data:    data,
	}
}
