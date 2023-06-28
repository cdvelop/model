package model

import (
	"strconv"
	"strings"
)

func (o *Object) BuildHtmlForm() string {

	for _, m := range o.Modules {

		if m != nil && m.Theme != nil && strings.Contains(m.MainName, o.Name) {
			var input_tags string

			for index, f := range o.RenderFields() {

				id := m.InputIdTemplate(o.Name, f.Name, strconv.Itoa(index))

				tag := f.Input.HtmlTag(id, f.Name, f.SkipCompletionAllowed)

				if f.Input.Name != "hidden" {

					input_tags += m.InputTemplate(o.Name, f.Name, f.Legend, f.Input.HtmlName(), tag, index) + "\n"

				} else {

					input_tags += tag
				}

			}

			return m.FormTemplate(o.Name, input_tags)
		}

	}

	return ""
}

func (o *Object) BuildJSInputFormModule() string {

	for _, m := range o.Modules {

		if m != nil && m.Theme != nil && strings.Contains(m.MainName, o.Name) {

			out_vars := m.JsFormVariablesTemplate(o.Name)

			for _, f := range o.RenderFields() {
				out_vars += m.JsInputVariableTemplate(f.Name, f.Input.HtmlName())
			}

			return out_vars

		}
	}

	return ""
}
