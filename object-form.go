package model

// setear valores del formulario
func (o *Object) ResetFormValues(form_jsValue any, reset_input_view bool) (err string) {
	// o.Log("ok form data:", o.ObjectName, o.FormData)

	if o.FormData != nil {
		for _, field := range o.Fields {

			if _, exist := o.FormData[field.Name]; exist {

				if !field.NotClearValueOnFormReset {

					// seteamos el valor a ""
					o.FormData[field.Name] = ""

					if reset_input_view {

						if field.Input != nil {
							// o.Log("ok reset", field.Name, field.InputName)
							if field.Input.ResetParameters != nil {

								field.Input.ResetParameters.Params["form"] = form_jsValue

								field.Input.ResetParameters.Params["field_name"] = field.Name

								o.Log(field.Input.CallWithEnableAndQueryParams(o))

							}
						}

					}
				}
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
