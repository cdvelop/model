package model

import (
	"strings"
)

func (o Object) verificationAllFields(data map[string]string) (wrongFields []string) {
	for _, field := range o.Fields {

		if !field.SkipCompletionAllowed { // campo requerido
			// total_required++
			if dataIn, exists := data[field.Name]; exists {

				if !field.Input.ValidateField(dataIn, field.SkipValidation) {
					wrongFields = append(wrongFields, errorMessage(dataIn, &field))
				}

			} else {
				// si no existe y es requerido y no es llave primaria  ni llave foránea
				// fmt.Printf("NO EXISTE %v DATO\n", field.Name)
				if !strings.Contains(field.Name, "id_"+o.Name) {
					// fmt.Printf("NO ES LLAVE PRIMARIA %v \n", table_name)
					wrongFields = append(wrongFields, errorMessage(dataIn, &field))
				}
			}

		} else { //campo no requerido pero si no viene vació verificar solo si lo requiere
			if dataIn, exists := data[field.Name]; exists && dataIn != "" {
				// total_required++
				if !field.Input.ValidateField(dataIn, field.SkipValidation) {
					wrongFields = append(wrongFields, errorMessage(dataIn, &field))
				}
			}
		}
	}

	return
}
