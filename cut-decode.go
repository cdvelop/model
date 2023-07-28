package model

import (
	"fmt"
)

func (o Object) DataDecode(cut_data ...CutData) ([]map[string]string, error) {
	var decodedData []map[string]string

	for _, c := range cut_data {

		if len(c.Index) != len(c.Data) {
			return nil, fmt.Errorf("tamaño index %d en DataDecode es diferente a la data enviada: %d", len(c.Index), len(c.Data))
		}

		if len(c.Index) > len(o.Fields) {
			return nil, fmt.Errorf("tamaño index: %d en DataDecode es mayor al del modelo del objeto %v: %d", len(c.Index), o.Name, len(c.Data))
		}

		data_decoded := make(map[string]string)

		for field_index, field := range o.Fields {

			if value_position, exist := c.Index[uint8(field_index)]; exist {

				value := c.Data[value_position]

				data_decoded[field.Name] = value

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
