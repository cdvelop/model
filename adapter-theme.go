package model

type Theme interface {
	// ModuleHtmlTemplate ej: nombre del modulo html y el contenido
	//<div id="` + module_name + `" class="slider_panel">` + content + `</div>
	ModuleHtmlTemplate(module_name, content string) string
	// MenuButtonTemplate ej: <li class="navbar-item"><a href="#` + module_name + `" tabindex="` + index + `" class="navbar-link" name="` + module_name + `">
	// <svg aria-hidden="true" focusable="false" class="fa-primary"><use xlink:href="#` + icon_id + `" /></svg>
	// <span class="link-text">` + title + `</span></a></li>
	MenuButtonTemplate(module_name, index, icon_id, title string) string

	FormTemplate(object_name, input_tags string) string
	InputTemplate(object_name, field_name, legend, html_name, input_tag string, index int) string
	InputIdTemplate(object_name, field_name, index string) string
	JsFormVariablesTemplate(object_name string) string
	JsInputVariableTemplate(field_name, html_name string) string
}
