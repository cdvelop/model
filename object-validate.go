package model

import "fmt"

// validar objeto
func (o Object) ValidateData(validate_all_fields bool, all_data ...*map[string]string) error {

	for _, data := range all_data {

		var wrongFields []string

		//verificar si campos pertenecen a la tabla
		for field_name := range *data {
			if _, exist := o.FieldExist(field_name); !exist {
				wrongFields = append(wrongFields, fmt.Sprintf("%v no pertenece a la tabla %v", field_name, o.Name))
			}
		}
		if len(wrongFields) != 0 {
			return wrongMessage(wrongFields)
		}

		if validate_all_fields {
			wrongFields = o.verificationAllFields(data)
		} else {

			// si no tiene su llave primaria valida..bye

			pk_value, exist := (*data)[o.PrimaryKeyName()]
			if !exist {
				return fmt.Errorf("Llave Primaria inexistente")
			}

			field := o.GetFieldByName(o.PrimaryKeyName())

			if !field.Input.ValidateField(pk_value, false) {
				return fmt.Errorf("Llave Primaria invalida")
			}

			wrongFields = o.verificationUpdateFields(data)
		}

		if len(wrongFields) != 0 {
			return wrongMessage(wrongFields)
		}
	}

	return nil
}
