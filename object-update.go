package model

func (o Object) verificationUpdateFields(data map[string]string) (wrongFields []string) {

	for field_name, value := range data {

		if field_found, exist := o.FieldExist(field_name); exist { // existe

			if !field_found.Inalterable {

				if !field_found.SkipCompletionAllowed || value != "" { // si es campo requerido se valida o distinto de vació

					if !field_found.Input.ValidateField(value, field_found.SkipValidation) {
						wrongFields = append(wrongFields, "dato: "+value+" en: "+field_found.Legend)
					}
				}

			} else {
				wrongFields = append(wrongFields, field_found.Legend+" no se puede modificar")
			}

		} else {
			wrongFields = append(wrongFields, "campo: "+field_name+" no existe en el modelo de la tabla")
		}
	}

	return
}
