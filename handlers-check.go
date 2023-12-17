package model

import (
	"reflect"
)

func (h MainHandler) CheckInterfaces(pkg_name string, struct_in interface{}) (err string) {
	const e = "CheckInterfaces "
	nils := h.getNilInterfaces()
	// fmt.Println("NILS:", nils)
	names := GetFieldNamesFrom(struct_in)

	// fmt.Println("NAMES:", names)
	for _, name := range names {
		for _, n := range nils {
			if name == n {
				return e + "en " + pkg_name + "handler nil: " + name
			}
		}
	}

	return ""
}

func (h MainHandler) getNilInterfaces() []string {
	var nullFields []string
	// Obtener el valor reflect.Value de MainHandler
	handlersValue := reflect.ValueOf(h)
	// Iterar sobre los campos de MainHandler
	for i := 0; i < handlersValue.NumField(); i++ {
		fieldValue := handlersValue.Field(i)
		// Verificar si el campo es una interfaz es nula
		if fieldValue.Kind() == reflect.Interface && fieldValue.IsNil() {
			fieldName := handlersValue.Type().Field(i).Name
			nullFields = append(nullFields, fieldName)
		}
	}

	return nullFields
}

func GetFieldNamesFrom(struct_in interface{}) (field_names []string) {
	// Obtener el tipo reflect.Type de la estructura
	structType := reflect.TypeOf(struct_in)
	// Iterar sobre los campos de la estructura
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		field_names = append(field_names, field.Name)
	}

	return field_names
}
