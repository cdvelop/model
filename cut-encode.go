package model

// DataEncode quita los nombres de los campos de la data según modelo del objeto
func (o Object) DataEncode(data map[string]string) CutData {

	cut_data := CutData{
		Index: make(map[uint8]uint8),
		Data:  []string{},
	}

	for field_index, f := range o.Fields {
		if value, exist := data[f.Name]; exist && value != "" {

			cut_data.Data = append(cut_data.Data, value)
			// indice del campo original + posición en el array de envió de la data
			cut_data.Index[uint8(field_index)] = uint8(len(cut_data.Data) - 1)
		} else {
			// eliminamos el campo del objeto original
			delete(data, f.Name)
		}
	}

	return cut_data
}
