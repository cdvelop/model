package model

type TestAdapter interface {
	AddUsesCaseTest(uses_case ...Response)
	//ej: os.Args found == "test"
	LoadE2Etests(osArgs []string) (e2eTests string)
	RunE2EfrontTests()

	TestActions() (BackEndActions, FrontEndActions)
}

type FrontEndActions struct {
	// acciones en el frontend
	Frontend_action string
	// objeto a usar para la acción
	Name_object_use string `Legend:"name"`
	// click id de un objeto especifico
	Clicking_ID string `Legend:"click"` // ej Click
	// click en le menu del modulo
	Click_menu_module string `Legend:"click"`
	// click en elementos del modulo del objeto ej: "button[name='btn_cancel']"
	Click_object_element string `Legend:"click"`
	//Completar formulario
	Form_complete string `Legend:"complete"`
	// añadir data a la existente
	Form_existing_add string `Legend:"add data"`
	//Espera en milisegundos ej 200,2000 = 2 seg
	Wait string `Legend:"wai"`
}

type BackEndActions struct {
	// acciones en el backend
	Backend_action string
	// setear la fecha del servidor a un dia especifico ej: 2023-12-24
	Set_server_date string
}
