package model

type Object struct {
	// nombre del componente u objeto ej: client, search_footer,datalist,form
	Name string

	TextFieldNames []string //nombre de campos mas representativos del objeto o tabla ej: name, address, phone
	Fields         []Field  //campos del objeto

	module *Module // módulo origen del objeto

	BackendRequest

	FrontendResponse
}

func (o Object) MainName() string {
	return o.Name
}

func (o *Object) Response(action, message string, data ...map[string]string) Response {

	return Response{
		Action:  action,
		Object:  o.Name,
		Message: message,
		Data:    data,
	}
}
