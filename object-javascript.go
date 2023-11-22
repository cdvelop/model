package model

func (o Object) CallJsFunctionObject(func_name string, enable bool, params ...map[string]interface{}) {
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
	o.CallFunction(func_name, all_params)

}

func (o Object) ClickElement(query_element string) error {
	return o.ElementClicking(o.QuerySelectorModule(o.ModuleName) + " " + query_element)
}

func (o Object) ClickModule() error {
	return o.ElementClicking(o.QuerySelectorMenuModule(o.ModuleName))
}

func (o Object) ClickingID(id string) error {

	module_html, err := o.GetHtmlModule(o.ModuleName)
	if err != nil {
		return err
	}

	if o.FrontHandler.ViewAdapter == nil {
		return Error("error objeto", o.ObjectName, "no tiene controlador ViewAdapter para realizar click")
	}

	o.CallFunction(o.FrontHandler.NameViewAdapter()+"ObjectClicking", module_html, id)

	return nil
}
