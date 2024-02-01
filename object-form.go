package model

func (o *Object) DeleteFormData() {
	if o.FormData != nil {
		for _, field := range o.Fields {

			if _, exist := o.FormData[field.Name]; exist {

				if !field.NotClearValueOnFormReset {
					// seteamos el valor a ""
					o.FormData[field.Name] = ""
				}
			}
		}
	}
}

// resetear la vista de los inputs del formulario
func (o *Object) ResetInputsViewForm(form_jsValue any) (err string) {
	// o.Log("ok form data:", o.ObjectName, o.FormData)

	for _, field := range o.RenderFields() {
		// o.Log("NOMBRE CAMPO exist:", field.Name)

		if field.Input != nil && field.Input.ResetParameters != nil {
			// o.Log("ok reset", field.Name, field.InputName)

			field.Input.ResetParameters.Params["form"] = form_jsValue

			field.Input.ResetParameters.Params["field_name"] = field.Name

			// Añadir los parámetros específicos de la función
			field.Input.ResetParameters.Params["enable"] = field.Input.ResetParameters.Enable

			if !field.Input.ResetParameters.NotSendQueryObject {
				field.Input.ResetParameters.Params["query"] = o.QuerySelectorObject(o.ModuleName, o.ObjectName)
			}

			_, err := o.CallFunction(field.Input.ResetParameters.ResetJsFuncName, struct {
				Params       map[string]any
				ResultInt    bool
				ResultString bool
			}{
				Params: field.Input.ResetParameters.Params,
			})
			if err != "" {
				o.Log("ResetInputsViewForm", field.Name, "input:", field.Input.InputName, "error:", err)
			}
		}

	}

	return
}

func (o Object) FieldsToFormValidate() (fiels_out []Field) {
	for _, field := range o.Fields {
		if !field.NotRenderHtml && field.Input != nil && field.SourceTable == "" {
			fiels_out = append(fiels_out, field)
		}
	}
	return
}

func (o Object) RequiredFields() (fiels_out []Field) {
	for _, field := range o.Fields {
		if !field.SkipCompletionAllowed {
			fiels_out = append(fiels_out, field)
		}
	}
	return
}
