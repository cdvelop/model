package model

import "fmt"

// validar objeto
func (o Object) ValidateData(validate_all_fields bool, all_data ...map[string]string) (string, bool) {

	for _, data := range all_data {

		var wrongFields []string

		//verificar si campos pertenecen a la tabla
		for field_name := range data {
			if _, exist := o.FieldExist(field_name); !exist {
				wrongFields = append(wrongFields, fmt.Sprintf("%v no pertenece a la tabla %v", field_name, o.Name))
			}
		}
		if len(wrongFields) != 0 {
			return wrongMessage(wrongFields), false
		}

		if validate_all_fields {
			wrongFields = o.verificationAllFields(data)
		} else {

			// si no tiene su llave primaria valida..bye

			pk_value, exist := data[o.PrimaryKeyName()]
			if !exist {
				return "Llave Primaria inexistente", false
			}

			field := o.FilterField(o.PrimaryKeyName())

			if !field.Input.ValidateField(pk_value, false) {
				return "Llave Primaria invalida", false
			}

			wrongFields = o.verificationUpdateFields(data)
		}

		if len(wrongFields) != 0 {
			return wrongMessage(wrongFields), false
		}
	}

	return "", true
}
