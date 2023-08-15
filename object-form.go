package model

import (
	"strconv"
	"strings"
)

func (o *Object) BuildHtmlForm(t Theme) string {

	if o.module != nil && t != nil && len(o.Fields) != 0 && strings.Contains(o.module.ModuleName, o.Name) {
		var input_tags string

		for index, f := range o.RenderFields() {

			id := t.InputIdTemplate(o.Name, f.Name, strconv.Itoa(index))

			tag := f.Input.HtmlTag(id, f.Name, f.SkipCompletionAllowed)

			if f.Input.InputName != "hidden" {

				input_tags += t.InputTemplate(o.Name, f.Name, f.Legend, f.Input.HtmlName(), tag, index) + "\n"

			} else {

				input_tags += tag
			}

		}

		return t.FormTemplate(o.Name, input_tags)
	}

	return ""
}

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
