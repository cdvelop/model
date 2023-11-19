package model

import (
	"strconv"
)

func (o Object) DataDecode(cut_data ...CutData) ([]map[string]string, error) {
	var decodedData []map[string]string

	for _, c := range cut_data {

		data_decoded := make(map[string]string)

		if pst_1, data_equals_index := c.Index[0]; data_equals_index && pst_1 == -1 {

			// fmt.Println("DATA IGUAL INDICE (-1)", pst_1, c.Data)

			if len(c.Data) != len(o.Fields) {
				return nil, Error("error tamaño data: " + strconv.Itoa(len(c.Data)) + " en DataDecode es diferente al del modelo del objeto " + o.ObjectName + ": " + strconv.Itoa(len(o.Fields)))
			}

			for field_index, field := range o.Fields {
				data_decoded[field.Name] = c.Data[field_index]
			}

		} else {
			// fmt.Println("DATA DIFERENTE INDICE", c.Data)

			if len(c.Index) != len(c.Data) {
				// return nil, fmt.Errorf("tamaño index %d en DataDecode es diferente a la data enviada: %d", len(c.Index), len(c.Data))
				return nil, Error("tamaño index " + strconv.Itoa(len(c.Index)) + " en DataDecode es diferente a la data enviada: " + strconv.Itoa(len(c.Data)))
			}

			if len(c.Index) > len(o.Fields) {
				return nil, Error("tamaño index: " + strconv.Itoa(len(c.Index)) + " en DataDecode es mayor al del modelo del objeto " + o.ObjectName + ": " + strconv.Itoa(len(o.Fields)))
			}

			for field_index, field := range o.Fields {

				if value_position, exist := c.Index[int8(field_index)]; exist {

					value := c.Data[value_position]

					data_decoded[field.Name] = value

				}
			}
		}

		decodedData = append(decodedData, data_decoded)
	}

	return decodedData, nil
}

func (cr *CutResponse) CutResponseDecode(data []map[string]string) (out Response) {

	if len(cr.CutOptions) > 0 && cr.CutOptions[0] != "" { //Action
		out.Action = cr.CutOptions[0]

	} else {
		out.Action = "error"
	}

	if len(cr.CutOptions) > 1 && cr.CutOptions[1] != "" { // Object
		// fmt.Println("si contiene objeto")
		out.Object = cr.CutOptions[1]

	} else {
		out.Object = "error"
	}

	if len(cr.CutOptions) > 2 && cr.CutOptions[2] != "" { //Message
		// fmt.Println("contiene mensaje")
		out.Message = cr.CutOptions[2]
	}

	out.Data = data

	return

}
