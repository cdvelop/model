package model

type Object struct {
	ApiHandler

	TextFieldNames []string //nombre de campos mas representativos del objeto o tabla ej: name, address, phone
	Fields         []Field  //campos del objeto

	*Module // módulo origen del objeto
	// Css() string ej: .mi-style{background:black;}
	Css

	JsGlobal
	// JsFunctions() string | ej: MyFunction(e){console.log("hello ",e)};
	JsFunctions
	// JsListeners() string | ej: btn.addEventListener('click', MyFunction);
	JsListeners

	// FolderPath() string ej: func (Object) FolderPath() string {
	// _, filename, _, _ := runtime.Caller(0)
	// dir := filepath.Dir(filename)
	// return filepath.ToSlash(dir)
	// }
	Path
}
