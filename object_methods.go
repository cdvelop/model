package model

import "strings"

func (o *Object) AddModule(m *Module, api_name ...string) {
	if m != nil {
		var (
			endpoint    string
			middle_name string
		)

		for i, v := range api_name {

			if v != o.Name {
				if i == 0 {
					endpoint = "." + v
				} else {
					endpoint += "." + v
				}
			}

		}

		if m.ModuleName != o.Name {
			middle_name = "." + o.Name
		}

		// change to unique api name
		o.apiName = m.ModuleName + middle_name + endpoint
		o.module = m
		m.Objects = append(m.Objects, o)
	}
}

func (o Object) Module() *Module {
	return o.module
}

func (o Object) Api() string {
	return o.apiName
}

func (o Object) ModuleName() string {
	return o.module.ModuleName
}

func (o Object) GetRepresentativeTextField(data_element map[string]string) (values string) {
	for _, keyNameModel := range o.TextFieldNames {
		if valueFound, ok := data_element[keyNameModel]; ok {
			values += valueFound + ` `
		}
	}
	values = strings.TrimSpace(values)
	return
}

func (o Object) Columns() (columns []string) {
	for _, f := range o.Fields {
		columns = append(columns, f.Name)
	}
	return
}

func (o Object) GetFieldByName(nameRq string) (fielOut Field) {
	for _, field := range o.Fields {
		if nameRq == field.Name {
			return field
		}
	}
	return
}

func (o Object) FilterFields(namesRequired ...string) (fielsOut []Field) {
	for _, field := range o.Fields {
		for _, nameRq := range namesRequired {
			if nameRq == field.Name {
				fielsOut = append(fielsOut, field)
				break
			}
		}
	}
	return
}

func (o Object) RenderFields() (fielsOut []Field) {
	for _, field := range o.Fields {
		if !field.NotRenderHtml {
			fielsOut = append(fielsOut, field)
		}
	}
	return
}

func (o Object) RequiredFields() (fielsOut []Field) {
	for _, field := range o.Fields {
		if !field.SkipCompletionAllowed {
			fielsOut = append(fielsOut, field)
		}
	}
	return
}

func (o Object) FilterRemoveFields(namesToRemove ...string) (fielsOut []Field) {
	removeNames := make(map[string]bool, len(namesToRemove))
	for _, name := range namesToRemove {
		removeNames[name] = false
	}

	for _, field := range o.Fields {
		if _, exist := removeNames[field.Name]; !exist {
			fielsOut = append(fielsOut, field)
		}
	}
	return
}

func (o Object) PrimaryKeyName() string {
	return "id_" + o.Name
}

func (o Object) GetID(data_search map[string]string) string {

	if id, found := data_search[o.PrimaryKeyName()]; found {
		return id
	}

	return "ERROR_NO_ID_" + o.Name
}

func (o Object) FieldExist(field_name string) (Field, bool) {
	for _, field := range o.Fields {
		if field.Name == field_name {
			return field, true
		}
	}
	return Field{}, false
}
