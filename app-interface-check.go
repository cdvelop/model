package model

import (
	"reflect"
)

func (h Handlers) CheckInterfaces(pkg_name string, struct_in interface{}) error {
	nils := h.getNilInterfaces()
	// fmt.Println("NILS:", nils)
	names := GetFieldNamesFrom(struct_in)
	// fmt.Println("NAMES:", names)
	for _, name := range names {
		for _, n := range nils {
			if name == n {
				return Error("error en", pkg_name, "handler nil:", name)
			}
		}
	}

	return nil
}

func (h Handlers) getNilInterfaces() []string {
	var nullFields []string
	// Obtener el valor reflect.Value de Handlers
	handlersValue := reflect.ValueOf(h)
	// Iterar sobre los campos de Handlers
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

func GetFieldNamesFrom(struct_in interface{}) []string {
	var fieldNames []string
	// Obtener el tipo reflect.Type de la estructura
	dataType := reflect.TypeOf(struct_in)
	// Iterar sobre los campos de la estructura
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		fieldNames = append(fieldNames, field.Name)
	}

	return fieldNames
}
