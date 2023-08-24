package model

import "strconv"

func (o *Object) BuildJSInputFormModule(t Theme) string {

	if t != nil && len(o.Fields) != 0 {

		out_vars := t.JsFormVariablesTemplate(o.Name)

		for _, f := range o.RenderFields() {
			out_vars += t.JsInputVariableTemplate(f.Name, f.Input.HtmlName())
		}

		return out_vars

	}

	return ""
}

func (o Object) TestData(required_objects int, skip_id, wrong_data bool) ([]map[string]string, error) {

	var out = make([]map[string]string, required_objects)

	for _, f := range o.Fields {

		if skip_id {
			if f.IsPrimaryKey(&o) {
				continue
			}
		}

		var obtained_data []string

		if !wrong_data {

			obtained_data = f.Input.GoodTestData()

		} else {

			obtained_data = f.Input.WrongTestData()

		}

		if required_objects > len(obtained_data) {
			return nil, MyError{
				Message: "error solicitud data fuera de rango " + f.Name + " tipo " + f.InputName + " solo contiene: " + strconv.Itoa(len(obtained_data))}
		}

		for i := 0; i < required_objects; i++ {

			previous := out[i]

			if previous == nil {
				previous = make(map[string]string) // Crear un nuevo mapa en cada iteración
			}

			previous[f.Name] = obtained_data[i]

			out[i] = previous

		}
	}

	return out, nil
}
