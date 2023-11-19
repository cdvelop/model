package model

import (
	"reflect"
)

// func (h Handlers) CheckInterfaces(pkg_name string, struct_in interface{}) error {
// 	nils := h.getNilInterfaces()
// 	names := GetFieldNamesFrom(struct_in)

// 	// Obtener el valor reflect.Value de struct_in
// 	structValue := reflect.ValueOf(struct_in)
// 	if structValue.Kind() == reflect.Ptr {
// 		// Si struct_in es un puntero, obtener el valor apuntado
// 		structValue = structValue.Elem()
// 	}

// 	for _, name := range names {
// 		for _, n := range nils {
// 			if name == n {
// 				return Error("error en " + pkg_name + " handler nil: " + name)
// 			}
// 		}
// 	}

// 	return nil
// }

func (h Handlers) CheckInterfaces(pkg_name string, struct_in interface{}) error {
	nils := h.getNilInterfaces()
	// fmt.Println("NILS:", nils)
	names, err := GetFieldNamesFrom(struct_in)
	if err != nil {
		return err
	}
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

func GetFieldNamesFrom(struct_in interface{}) ([]string, error) {
	var fieldNames []string
	// Obtener el tipo reflect.Type de la estructura
	structType := reflect.TypeOf(struct_in)
	// Iterar sobre los campos de la estructura
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldNames = append(fieldNames, field.Name)
	}

	return fieldNames, nil
}

func GetFieldNamesFromNEW(struct_in interface{}) ([]string, error) {
	var fieldNames []string
	// Obtener el tipo reflect.Type de la estructura
	structType := reflect.TypeOf(struct_in)
	valueOfStruct := reflect.ValueOf(struct_in)
	if structType.Kind() == reflect.Ptr {
		if valueOfStruct.IsNil() {
			return nil, Error("GetFieldNamesFrom error. el puntero de la estructura no puede ser nulo")
		}
		// Si struct_in es un puntero, obtener el tipo apuntado
		structType = structType.Elem()
	}
	// Iterar sobre los campos de la estructura
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldNames = append(fieldNames, field.Name)
	}

	return fieldNames, nil
}
