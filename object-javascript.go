package model

func (o Object) ClickObjectElement(query_element string) (err string) {
	return o.ElementClicking(o.QuerySelectorModule(o.ModuleName) + " " + query_element)
}

func (o Object) ClickMenuModule() (err string) {
	return o.ElementClicking(o.QuerySelectorMenuModule(o.ModuleName))
}

// si no se proporciona id_param se buscara en el formulario actual del objeto
func (o Object) ClickingID(id_param ...string) (err string) {
	const e = ". ClickingID"
	var id string

	for _, v := range id_param {
		id = v
	}

	if id == "" { // no se proporciono id lo buscamos en el formulario
		id, _ = o.GetID() // este función retorna error pero podemos hacer click en los elementos sin id
	}

	module_html, err := o.CheckModuleHtml()
	if err != "" {
		return err + e
	}

	_, err = o.CallFunction(o.FrontHandler.ViewHandlerName()+"ObjectClicking", struct {
		Params       map[string]any
		ResultInt    bool
		ResultString bool
	}{
		Params: map[string]any{"module": module_html, "id": id},
	})
	if err != "" {
		return err + e
	}

	return
}

func (o Object) CountViewElements() (total int, err string) {
	const t = "CountViewElements "

	js_value, err := o.CallFunction(o.FrontHandler.ViewHandlerName()+"ObjectCount", struct {
		Params       map[string]any
		ResultInt    bool
		ResultString bool
	}{
		ResultInt: true,
	})
	if err != "" {
		return 0, t + err
	}

	var ok bool
	total, ok = js_value.(int)
	if !ok {
		return 0, t + "valor retornado en la función " + o.FrontHandler.ViewHandlerName() + "ObjectCount" + " no es de tipo int"
	}

	return
}

func (o Object) CheckModuleHtml() (module_html any, err string) {

	if o.FrontHandler.ViewHandlerObject == nil {
		return nil, "error objeto " + o.ObjectName + " no tiene controlador ViewHandlerObject"
	}

	module_html, err = o.GetHtmlModuleContent()
	if err != "" {
		return nil, err
	}

	return
}

// select_items ej: div#user, `button[name="login"]`
func (o Object) GetHtmlModuleContent(select_items ...string) (jsValue any, err string) {
	var selected string

	for _, item := range select_items {
		selected += " " + item
	}

	jsValue, err = o.SelectContent(struct {
		QuerySelector string
		GetContent    string
		SetAfter      bool
		StringReturn  bool
	}{
		QuerySelector: o.QuerySelectorModule(o.ModuleName) + selected,
	})

	if err != "" {
		err = "GetHtmlModuleContent " + err
	}

	return
}
