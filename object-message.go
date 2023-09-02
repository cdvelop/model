package model

import (
	"strings"
)

func wrongMessage(wrongFields []string) error {

	// log.Println("WRONG FIELDS: *", wrongFields, "*")

	defMessage := `campo incorrecto`
	if len(wrongFields) > 1 {
		defMessage = `campos incorrectos`
	}

	wf := strings.Join(wrongFields, ", ")

	// fmt.Printf("WF %v\n", wrongFields)

	// return fmt.Errorf("%v. %v", defMessage, strings.TrimSpace(wf))

	return Error(defMessage + ". " + strings.TrimSpace(wf))
}

func errorMessage(dataIn string, f *Field, err error) (out string) {

	if err != nil {
		out = err.Error() + " "
	}

	var wrong_field string

	if f.Legend != "" {
		wrong_field = f.Legend
	} else {
		wrong_field = f.Name
	}

	if dataIn == "" {
		out += "[" + wrong_field + "] dato requerido"
	} else {
		out += "dato: [" + dataIn + "] en: [" + wrong_field + "]"
	}

	return
}
