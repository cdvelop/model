package model

import "strconv"

// type ModuleTestAdapter interface {
// 	RunModuleTests(t ...Response) (err string)
// }

type Tests struct {
	*Object
	NameObject string `Legend:"Nombre Objeto"`

	ClickingID   string `Legend:"Click en Objeto"` // ej Click
	ClickModule  string `Legend:"Click en Modulo"`
	ClickElement string `Legend:"Click en Elemento html"` // ej: "button[name='btn_cancel']"

	Wait         string `Legend:"Espera en milisegundos"` // ej 200
	milliseconds int

	all_test []Response
}

func (a *Tests) RunModuleTests(all_tests ...Response) (err string) {

	a.milliseconds = 100

	for _, tests := range all_tests {
		for i, data := range tests.Data {
			a.processData(i, data, func(index int, er string) {

				a.Log("INDICE DATA ACTUAL:", index, "error:", err)
			})
		}
	}

	return ""
}

func (a *Tests) processData(i int, d map[string]string, callback func(index int, er string)) {

	index := i

	if d[a.Wait] != "" {

		wait, err := strconv.Atoi(d[a.Wait])
		if err == nil {
			a.milliseconds += wait
		}
	}
	if d[a.ClickModule] != "" {
		a.WaitFor(a.milliseconds, func() {
			a.Log("CLICK MODULO:", d[a.ModuleName])
			object, err := a.GetObjectByName(d[a.NameObject])
			if err != nil {
				callback(index, err.Error())
				return
			}

			err = object.ClickModule()
			if err != nil {
				callback(index, err.Error())
				return
			}

		})
	}
	if d[a.ClickingID] != "" {
		a.WaitFor(a.milliseconds, func() {
			a.Log("CLICK OBJETO ID:", d[a.ClickingID])
			object, err := a.GetObjectByName(d[a.NameObject])
			if err != nil {
				callback(index, err.Error())
				return
			}

			a.Log("CLICK EN", object.ObjectName)
			err = object.ClickingID(d[a.ClickingID])
			if err != nil {
				callback(index, err.Error())
				return
			}
		})
	}
	if d[a.ClickElement] != "" {
		a.WaitFor(a.milliseconds, func() {
			a.Log("ClickElement:", d[a.ClickElement])
			object, err := a.GetObjectByName(d[a.NameObject])
			if err != nil {
				callback(index, err.Error())
				return
			}

			err = object.ClickElement(d[a.ClickElement])
			if err != nil {
				callback(index, err.Error())
				return
			}
		})
	}

	// a.remaining_data--

	callback(index, "")
}
