package model

import (
	"strconv"
	"strings"
)

func (o *Object) BuildHtmlForm() string {

	if o.Module != nil && o.Module.Theme != nil && len(o.Fields) != 0 && strings.Contains(o.Module.MainName, o.Name) {
		var input_tags string

		for index, f := range o.RenderFields() {

			id := o.Module.InputIdTemplate(o.Name, f.Name, strconv.Itoa(index))

			tag := f.Input.HtmlTag(id, f.Name, f.SkipCompletionAllowed)

			if f.Input.Name != "hidden" {

				input_tags += o.Module.InputTemplate(o.Name, f.Name, f.Legend, f.Input.HtmlName(), tag, index) + "\n"

			} else {

				input_tags += tag
			}

		}

		return o.Module.FormTemplate(o.Name, input_tags)
	}

	return ""
}

func (o *Object) BuildJSInputFormModule() string {

	if o.Module != nil && o.Module.Theme != nil && len(o.Fields) != 0 {

		out_vars := o.Module.JsFormVariablesTemplate(o.Name)

		for _, f := range o.RenderFields() {
			out_vars += o.Module.JsInputVariableTemplate(f.Name, f.Input.HtmlName())
		}

		return out_vars

	}

	return ""
}
