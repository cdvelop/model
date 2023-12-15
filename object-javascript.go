package model

func (o Object) ClickObjectElement(query_element string) (err string) {
	return o.ElementClicking(o.QuerySelectorModule(o.ModuleName) + " " + query_element)
}

func (o Object) ClickMenuModule() (err string) {
	return o.ElementClicking(o.QuerySelectorMenuModule(o.ModuleName))
}

func (o Object) ClickingID(id string) (err string) {

	module_html, err := o.CheckModuleHtml()
	if err != "" {
		return "ClickingID " + err + " para realizar click"
	}

	_, err = o.CallFunction(o.FrontHandler.NameViewAdapter()+"ObjectClicking", module_html, id)

	return
}

func (o Object) CountViewElements() (total int, err string) {
	const t = "CountViewElements "

	fn := CallJsOptions{
		NameJsFunc:      o.FrontHandler.NameViewAdapter() + "ObjectCount",
		Enable:          false,
		SendQueryObject: true,
		Params:          map[string]any{},
		ResultInt:       true,
		ResultString:    false,
	}

	js_value, err := fn.CallWithEnableAndQueryParams(&o)
	if err != "" {
		return 0, t + err
	}

	var ok bool
	total, ok = js_value.(int)
	if !ok {
		return 0, t + "valor retornado en la función " + fn.NameJsFunc + " no es de tipo int"
	}

	return
}

func (o Object) CheckModuleHtml() (module_html any, err string) {

	if o.FrontHandler.ObjectViewHandler == nil {
		return nil, "error objeto " + o.ObjectName + " no tiene controlador ObjectViewHandler"
	}

	module_html, err = o.GetHtmlModule(o.ModuleName)
	if err != "" {
		return nil, err
	}

	return
}
