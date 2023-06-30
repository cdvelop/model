package model

import "strings"

func (o Object) MainName() string {
	return o.Name
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

func (o Object) IsExternalComponent() bool {

	if o.Path != nil && o.Path.FolderPath() != "" {
		return true
	}

	return false
}
