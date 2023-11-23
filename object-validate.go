package model

import (
	"github.com/cdvelop/strings"
)

type AlternativeValidateAdapter interface {
	ValidateData(its_new, its_update_or_delete bool, all_data ...map[string]string) (err string)
}

// validar data objeto
func (o Object) ValidateData(its_new, its_update_or_delete bool, all_data ...map[string]string) (err string) {

	if o.AlternativeValidateAdapter != nil {
		return o.AlternativeValidateAdapter.ValidateData(its_new, its_update_or_delete, all_data...)
	}

	for _, data := range all_data {

		var wrongFields []string

		//verificar si campos pertenecen a la tabla
		for field_name := range data {
			if _, exist := o.FieldExist(field_name); !exist {
				wrongFields = append(wrongFields, field_name+" no pertenece a la tabla "+o.Table)
			}
		}
		if len(wrongFields) != 0 {
			return wrongMessage(wrongFields)
		}

		if its_new {
			wrongFields = o.verificationAllFields(data)
		} else {

			if its_update_or_delete {
				// si no tiene su llave primaria valida..???

				pk_value, exist := data[o.PrimaryKeyName()]
				if !exist {
					return "llave Primaria inexistente"
				}

				field, err := o.GetFieldsByNames(o.PrimaryKeyName())
				if err != "" {
					return err
				}

				err = field[0].Input.ValidateField(pk_value, false)
				if err != "" {
					return "llave Primaria invalida " + err
				}

			}

			wrongFields = o.verificationUpdateFields(data)
		}

		if len(wrongFields) != 0 {
			return wrongMessage(wrongFields)
		}
	}

	return ""
}

func (o Object) verificationAllFields(data map[string]string) (wrongFields []string) {
	for _, field := range o.Fields {

		if !field.SkipCompletionAllowed { // campo requerido
			// total_required++
			if dataIn, exists := data[field.Name]; exists {

				err := field.Input.ValidateField(dataIn, field.SkipCompletionAllowed)
				if err != "" {
					wrongFields = append(wrongFields, errorMessage(dataIn, &field, err))
				}

			} else {
				// si no existe y es requerido y no es llave primaria  ni llave foránea
				// fmt.Printf("NO EXISTE %v DATO\n", field.Name)
				if strings.Contains(field.Name, o.PrimaryKeyName()) == 0 {
					// fmt.Printf("NO ES LLAVE PRIMARIA %v \n", table_name)
					wrongFields = append(wrongFields, errorMessage(dataIn, &field, ""))
				}
			}

		} else { //campo no requerido pero si no viene vació verificar solo si lo requiere
			if dataIn, exists := data[field.Name]; exists && dataIn != "" {
				// total_required++
				err := field.Input.ValidateField(dataIn, field.SkipCompletionAllowed)
				if err != "" {
					wrongFields = append(wrongFields, errorMessage(dataIn, &field, err))
				}
			}
		}
	}

	return
}

func (o Object) verificationUpdateFields(data map[string]string) (wrongFields []string) {

	for field_name, value := range data {

		if field_found, exist := o.FieldExist(field_name); exist { // existe

			if !field_found.Unique {

				if !field_found.SkipCompletionAllowed || value != "" { // si es campo requerido se valida o distinto de vació

					err := field_found.Input.ValidateField(value, field_found.SkipCompletionAllowed)
					if err != "" {
						wrongFields = append(wrongFields, "dato: "+value+" en: "+field_found.Legend+" "+err)
					}
				}

			} else {
				if field_found.Legend == "" {
					field_found.Legend = field_found.Name
				}

				wrongFields = append(wrongFields, field_found.Legend+" es un campo único no se puede modificar")
			}

		} else {
			wrongFields = append(wrongFields, "campo: "+field_name+" no existe en el modelo de la tabla")
		}
	}

	return
}
