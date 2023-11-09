# Estructuras que utilizo con frecuencia en proyectos escritos en Go


## ejemplo de uso en un objeto
```go
func MedicalHistory() model.Object {
	return model.Object{
		Name:           "medicalhistory",
		PrincipalFieldsName: []string{"service_name"},
		Fields: []model.Field{
			{Name: "id_medicalhistory", Legend: "id", Input: unixid.InputPK()},
			{Name: "id_patient", Legend: "paciente", Input: unixid.InputPK()},

			{Name: "staff_id", Legend: "Id", Input: unixid.InputPK()},
			{Name: "staff_name", Legend: "Atendido por", Input: input.Text()},
			{Name: "staff_ocupation", Legend: "Especialidad", Input: input.Text()},
			{Name: "staff_area", Legend: "Area", Input: input.Check(Area{})},

			{Name: "service_id", Legend: "Prestación", Input: unixid.InputPK()},

			{Name: "day_attention", Legend: "Dia Atención", Input: input.Date()},
			{Name: "hour_attention", Legend: "Hora", Input: input.Hour(`min="08:15"`, `max="20:15"`)},

			{Name: "reason", Legend: "Motivo Consulta, Servicio o Procedimiento", Input: input.TextArea()},
			{Name: "diagnostic", Legend: "Diagnostico (Descripción)", Input: input.TextArea()},
			{Name: "prescription", Legend: "Prescripción (Conclusion)", Input: input.TextArea()},

			{Name: "last_print_file", Legend: "Ultima Impresión en Archivo", Input: input.Text()},
		},
	}
}
```