package model

// actualizar objeto en uso actualmente en el modulo
func (o *Object) SetActualObject() {
	o.Module.object_actual = o
}

func (o Object) Columns() (columns []string) {
	for _, f := range o.Fields {
		columns = append(columns, f.Name)
	}
	return
}

func (o Object) GetFieldsByNames(required_names ...string) (fiels_out []Field, err string) {

	if len(required_names) == 0 {
		return nil, "error debes ingresar mínimo un nombre de campo para buscar"
	}

	//NOTA: en webAssembly array de Puntero no funciona correctamente

	for _, field := range o.Fields {
		for _, nameRq := range required_names {
			if nameRq == field.Name {
				fiels_out = append(fiels_out, field)
				break
			}
		}
	}

	if len(fiels_out) != 0 {
		return fiels_out, ""
	}

	var space string
	var err_names string
	for _, name := range required_names {
		err_names += space + name
		space = " "
	}

	return nil, "campo(s) " + err_names + " no encontrado(s)"
}

func (o Object) RenderFields() (fiels_out []Field) {
	for _, field := range o.Fields {
		if !field.NotRenderHtml && field.Input != nil {
			if field.Input.HtmlName() != "hidden" {
				fiels_out = append(fiels_out, field)
			}
		}
	}
	return
}

func (o Object) FilterRemoveFields(namesToRemove ...string) (fiels_out []Field) {
	removeNames := make(map[string]bool, len(namesToRemove))
	for _, name := range namesToRemove {
		removeNames[name] = false
	}

	for _, field := range o.Fields {
		if _, exist := removeNames[field.Name]; !exist {
			fiels_out = append(fiels_out, field)
		}
	}
	return
}

func (o Object) GetID() (id, err string) {
	const e = ". GetID"

	err = o.FormCheck()
	if err != "" {
		return "", err + e
	}

	var ok bool
	if id, ok = o.FormData[o.PrimaryKeyName()]; !ok || id == "" {
		return "", "id no encontrado en el formulario del objeto:" + o.ObjectName + e
	}

	return id, ""
}

func (o Object) FormCheck() (err string) {
	const e = ". FormCheck"

	if o.FormData == nil {
		return "formulario del objeto:" + o.ObjectName + " no definido (nil)" + e
	}

	if len(o.FormData) == 0 {
		return "no hay elementos en el formulario del objeto:" + o.ObjectName + e
	}

	return
}

func (o Object) FieldExist(field_name string) (*Field, bool) {
	for _, field := range o.Fields {
		if field.Name == field_name {
			return &field, true
		}
	}
	return nil, false
}
