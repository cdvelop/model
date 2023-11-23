package model

// DataEncode quita los nombres de los campos de la data según modelo del objeto
func (o Object) DataEncode(all_data ...map[string]string) (out []CutData, err string) {

	if len(all_data) == 0 {
		return
	}

	for _, data := range all_data {

		// fmt.Println("ENTRO ACA:", data)

		var index_equal_positions = true

		cut_data := CutData{
			Index: make(map[int8]int8),
			Data:  []string{},
		}
		new_post := int8(0) // Inicializar el siguiente índice en 0

		for field_index, f := range o.Fields {
			if value, exist := data[f.Name]; exist && value != "" {

				new_idx := int8(field_index)

				// fmt.Println(f.Name, " new_idx:", new_idx, " new_post:", new_post)

				cut_data.Data = append(cut_data.Data, value)
				// indice del campo original + posición en el array de envió de la data
				cut_data.Index[new_idx] = new_post

				if new_idx != new_post {
					index_equal_positions = false
				}

				new_post++

			} else {
				// eliminamos el campo del objeto original
				delete(data, f.Name)

				index_equal_positions = false
			}
		}

		// fmt.Println("VALOR index_equal_positions", index_equal_positions)

		if index_equal_positions && len(o.Fields) == len(cut_data.Data) {
			cut_data.Index = map[int8]int8{0: -1} // todo igual a indice
			// fmt.Println("** DATA IGUAL = INDICE (-1)", data)
		}

		// fmt.Println("** CUT_DATA ", cut_data)
		if len(cut_data.Data) != 0 {
			out = append(out, cut_data)
		}
	}

	if len(out) == 0 {
		return nil, "error DataEncode campos enviados no coinciden con el modelo del objeto:" + o.ObjectName
	}

	return out, ""
}
