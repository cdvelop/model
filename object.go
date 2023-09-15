package model

type Object struct {
	// ej: client, search_footer,datalist
	Name string

	TextFieldNames []string //nombre de campos mas representativos ej: name, address, phone
	Fields         []Field  //campos

	*Module // módulo origen

	BackendHandler

	FrontendHandler
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
