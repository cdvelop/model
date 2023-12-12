package model

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
