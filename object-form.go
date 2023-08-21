package model

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
