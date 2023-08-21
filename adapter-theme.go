package model

type Theme interface {
	// MenuButtonTemplate ej: <li class="navbar-item"><a href="#` + module_name + `" tabindex="` + index + `" class="navbar-link" name="` + module_name + `">
	// <svg aria-hidden="true" focusable="false" class="fa-primary"><use xlink:href="#` + icon_id + `" /></svg>
	// <span class="link-text">` + title + `</span></a></li>
	MenuButtonTemplate(module_name, index, icon_id, title string) string
	MenuClassName() string // ej: .menu-container
	MenuItemClass() string // ej: navbar-item
	// FormTemplate(object_name, input_tags string) string
	// InputTemplate(object_name, field_name, legend, html_name, input_tag string, index int) string
	// InputIdTemplate(object_name, field_name, index string) string
	JsFormVariablesTemplate(object_name string) string
	JsInputVariableTemplate(field_name, html_name string) string

	ModuleClassName() string //ej: slider_panel
	ModuleTemplate(m *Module, form *Object) string
}
