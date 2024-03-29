package model

//ej: "id_client", options: "add_prefix_table_name" = "client.id_client"
func (o Object) PrimaryKeyName(options ...string) string {
	var prefix string
	for _, opt := range options {
		if opt == "add_prefix_table_name" {
			prefix = o.Table + "."
		}
	}

	return prefix + PREFIX_ID_NAME + o.Table
}

// ej: id_patient,patient_name,patient_run -
// add_prefix_table_name: true = patient.id_patient,patient.patient_name,patient.patient_run
func (o Object) JoinFieldNames(add_prefix_table_name bool) (out string) {
	var comma string

	for _, f := range o.Fields {
		if !f.NotRequiredInDB {
			// fmt.Println(o.Table, "CAMPO REQUERIDO EN DB:", f.Name)
			if add_prefix_table_name {
				if f.SourceTable != "" {
					out += comma + f.SourceTable + "." + f.Name
				} else {
					out += comma + o.Table + "." + f.Name
				}
			} else {
				out += comma + f.Name
			}

			comma = ","
		}
	}
	return
}

//solo campos requeridos en la base de datos. NOTA: puntero []*Field no funciona con slice
func (o Object) OnlyRequiredDbFieldsThisObject() (db_field []Field) {
	// fmt.Println("TABLA:", o.Table)
	for _, f := range o.Fields {
		if !f.NotRequiredInDB && f.SourceTable == "" {
			// fmt.Println("CAMPO:", f.Name, "SOURCE:", f.SourceTable)
			db_field = append(db_field, f)
		}
	}
	return
}

func (o Object) GetTablesNames() (out []string) {

	out = append(out, o.Table)

	for _, f := range o.Fields {

		if f.SourceTable != "" {

			for _, exist_table := range out {
				if exist_table != f.SourceTable {
					out = append(out, f.SourceTable)
					break
				}
			}
		}

	}

	return
}
