# Estructuras que utilizo con frecuencia en proyectos escritos en Go


## ejemplo de uso en una tabla
```go
func Patient() *model.Object {
	run_Global := input.RunGlobal{
		PlaceHolder: "ch ej: 13173472-7 dni ej: 12345678 ",
	}
	textOnly := input.TextOnly{}
	phoneNumber := input.Number{Pattern: `^[0-9]{7,11}$`}

	dateAge := input.DateAge{}

	t := model.Object{
		Noun:           "patient",
		TextFieldNames: []string{"patient_name"},
		Fields: []model.Field{
			
            {Name: "id_patient", DataType: "TEXT", Legend: "Id", Input: pkHiddenNoRequiredInView, NotRenderHtml: true},
			
            {Name: "patient_run", DataType: "TEXT", Unique: true, Legend: "Run o Dni", Input: run_Global},
			
            {Name: "patient_name", DataType: "TEXT", Legend: "Nombre y Apellido(s)", Input: textOnly},
			
            {Name: "patient_birthday", DataType: "TEXT", Legend: "Edad", Input: dateAge,},
			
            {Name: "patient_contact", DataType: "TEXT", Legend: "No Teléfono", Input: phoneNumber, SkipCompletionAllowed: true,Render: true},
			
            {Name: "patient_gender", DataType: "TEXT", Legend: "Genero", Input: radioGenero, SkipValidation: true},
			
            {Name: "patient_address", DataType: "TEXT", Legend: "Dirección", Input: textPointArea},
		},
	}

	return &t
}
```