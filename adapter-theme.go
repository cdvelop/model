package model

type Theme interface {
	// MenuButtonTemplate ej: <li class="navbar-item"><a href="#` + module_name + `" tabindex="` + index + `" class="navbar-link" name="` + module_name + `">
	// <svg aria-hidden="true" focusable="false" class="fa-primary"><use xlink:href="#` + icon_id + `" /></svg>
	// <span class="link-text">` + title + `</span></a></li>
	MenuButtonTemplate(module_name, index, icon_id, title string) string
	MenuClassName() string // ej: .menu-container
	MenuItemClass() string // ej: navbar-item

	ModuleClassName() string //ej: slider_panel
	ModuleTemplate(m *Module, form *Object, options ...string) string

	FunctionMessageName() string // ej: ShowMessageToUser
}
