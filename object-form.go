package model

// setear valores del formulario
func (o *Object) ResetFormValues(reset_input_view bool) (err string) {
	const this = "ResetFormValues error "
	if o.FormData == nil {
		return this + "data formulario es nil"
	}

	for _, field := range o.Fields {

		if _, exist := o.FormData[field.Name]; exist {

			if !field.NotClearValueOnFormReset {

				// seteamos el valor a ""
				o.FormData[field.Name] = ""

				if reset_input_view {
					if field.Input.ResetViewAdapter != nil {
						err = field.Input.ResetAdapterView()
						if err != "" {
							return this + field.Name + " " + err
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
