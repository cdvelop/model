package model

func (o Object) CallJsFunctionObject(func_name string, enable bool, params ...map[string]interface{}) (err string) {
	// Crear el mapa que contendrá todos los parámetros
	all_params := make(map[string]interface{})

	// Añadir los parámetros específicos de la función
	all_params["enable"] = enable
	all_params["query"] = o.QuerySelectorObject(o.ModuleName, o.ObjectName)

	// Añadir los parámetros adicionales proporcionados por el usuario
	for _, param := range params {
		for key, value := range param {
			all_params[key] = value
		}
	}

	// Llamar a la función JavaScript
	return o.CallFunction(func_name, all_params)
}

func (o Object) ClickObjectElement(query_element string) (err string) {
	return o.ElementClicking(o.QuerySelectorModule(o.ModuleName) + " " + query_element)
}

func (o Object) ClickMenuModule() (err string) {
	return o.ElementClicking(o.QuerySelectorMenuModule(o.ModuleName))
}

func (o Object) ClickingID(id string) (err string) {

	module_html, e := o.GetHtmlModule(o.ModuleName)
	if e != "" {
		return e
	}

	if o.FrontHandler.ObjectViewHandler == nil {
		return "error objeto " + o.ObjectName + " no tiene controlador ObjectViewHandler para realizar click"
	}

	return o.CallFunction(o.FrontHandler.NameViewAdapter()+"ObjectClicking", module_html, id)
}
