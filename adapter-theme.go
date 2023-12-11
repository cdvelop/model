package model

type ThemeAdapter interface {
	// MenuButtonTemplate ej: <li class="navbar-item"><a href="#` + module_name + `" tabindex="` + index + `" class="navbar-link" name="` + module_name + `">
	// <svg aria-hidden="true" focusable="false" class="fa-primary"><use xlink:href="#` + icon_id + `" /></svg>
	// <span class="link-text">` + title + `</span></a></li>
	MenuButtonTemplate(module_name, index, icon_id, title string) string
	MenuClassName() string // ej: .menu-container
	MenuItemClass() string // ej: navbar-item

	ModuleClassName() string //ej: slider_panel

	ModuleTemplate(*TemplateModuleConfig) string

	FunctionMessageName() string // ej: ShowMessageToUser
	// ej query: "div#"+o.ModuleName+" [data-id='"+o.ObjectName+"']"
	QuerySelectorMenuModule(module_name string) string
	QuerySelectorModule(module_name string) string
	QuerySelectorObject(module_name, object_name string) string
	QuerySelectorUserName() string
	QuerySelectorUserArea() string
}

type TemplateModuleConfig struct {
	Module *Module
	// ej:	`<div class="target-module">
	// 	<select name="select">
	// 		<option value="value1">Value 1</option>
	// 		<option value="value2" selected>Value 2</option>
	// 	</select>
	// </div>`
	HeaderInputTarget string
	Form              *Object
	AsideList         ContainerViewAdapter
	ButtonLogin       bool
	ButtonPrint       bool
}
