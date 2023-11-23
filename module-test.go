package model

import "strconv"

type Tests struct {
	*Object
	NameObject string `Legend:"Nombre Objeto"`

	ClickingID         string `Legend:"Click en Objeto"` // ej Click
	ClickMenuModule    string `Legend:"Click en Modulo"`
	ClickObjectElement string `Legend:"Click en Elemento html"` // ej: "button[name='btn_cancel']"

	Wait         string `Legend:"Espera en milisegundos"` // ej 200
	milliseconds int

	current_object *Object
}

func (a *Tests) RunModuleTests(all_tests []Response, result func(err string)) {

	var total_test = len(all_tests)
	var completedTests int
	for _, tests := range all_tests { // = all_tests []Response

		var total_data = len(tests.Data)
		var completedData int

		for index_data, data := range tests.Data { // tests.Data = Data []map[string]string

			a.processData(index_data, data, func(err string) {

				if err != "" {
					result(err)
					return
				}

				// Verificar si es la última iteración para pruebas y datos
				completedData++
				if completedTests == total_test && completedData == total_data {
					result("")
				}

			})
		}
		completedTests++
	}

}

func (a *Tests) processData(i int, d map[string]string, result func(err string)) {
	a.milliseconds = 100
	var err string

	if d[a.Wait] != "" {
		wait, e := strconv.Atoi(d[a.Wait])
		if e == nil {
			a.milliseconds += wait
		}
	}

	if d[a.NameObject] != "" {
		a.current_object, err = a.GetObjectByName(d[a.NameObject])
		if err != "" {
			result(err)
			return
		}
	}

	a.WaitFor(a.milliseconds, func() {

		if d[a.ClickMenuModule] != "" {
			a.Log("CLICK MODULO:", d[a.ModuleName])

			err = a.current_object.ClickMenuModule()
			if err != "" {
				result(err)
				return
			}

		}

		if d[a.ClickingID] != "" {
			a.Log("CLICK OBJETO ID:", d[a.ClickingID])

			err = a.current_object.ClickingID(d[a.ClickingID])
			if err != "" {
				result(err)
				return
			}
		}

		if d[a.ClickObjectElement] != "" {
			a.Log("ClickObjectElement:", d[a.ClickObjectElement])

			err = a.current_object.ClickObjectElement(d[a.ClickObjectElement])
			if err != "" {
				result(err)
				return
			}
		}

		result("")
	})

}
