package model

// ej: id_patient,patient_name,patient_run
func (o Object) JoinFieldNames() (out string) {
	var comma string
	for _, f := range o.Fields {
		out += comma + f.Name
		comma = ","
	}

	return
}
