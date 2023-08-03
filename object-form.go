package model

import (
	"strconv"
	"strings"
)

func (o *Object) BuildHtmlForm() string {

	if o.module != nil && o.module.Theme != nil && len(o.Fields) != 0 && strings.Contains(o.module.ModuleName, o.Name) {
		var input_tags string

		for index, f := range o.RenderFields() {

			id := o.module.InputIdTemplate(o.Name, f.Name, strconv.Itoa(index))

			tag := f.Input.HtmlTag(id, f.Name, f.SkipCompletionAllowed)

			if f.Input.InputName != "hidden" {

				input_tags += o.module.InputTemplate(o.Name, f.Name, f.Legend, f.Input.HtmlName(), tag, index) + "\n"

			} else {

				input_tags += tag
			}

		}

		return o.module.FormTemplate(o.Name, input_tags)
	}

	return ""
}

func (o *Object) BuildJSInputFormModule() string {

	if o.module.Theme != nil && len(o.Fields) != 0 {

		out_vars := o.module.JsFormVariablesTemplate(o.Name)

		for _, f := range o.RenderFields() {
			out_vars += o.module.JsInputVariableTemplate(f.Name, f.Input.HtmlName())
		}

		return out_vars

	}

	return ""
}
