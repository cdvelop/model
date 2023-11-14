package model

type IdHandler interface {
	GetNewID() string
}

type DataBaseAdapter interface {
	IdHandler
	RunOnClientDB() bool //verdadero corren en el cliente ej browser. por defecto falso corre en el servidor
	// items support in server db: []map[string]string, map[string]string
	CreateObjectsInDB(table_name string, backup_required bool, items any) error

	// from_tables ej: "users,products" or: public.reservation, public.patient"
	// params: map[string]string ej:{
	// LIMIT: 10, 5, 100. note: Postgres y MySQL: "LIMIT 10", SQLite: "LIMIT 10 OFFSET 0" OR "" no limit
	// ORDER_BY: name,phone,address
	// SELECT: "name, phone, address" default *
	// WHERE: "patient.id_patient = reservation.id_patient AND reservation.id_staff = '2'"
	// ARGS: "1,4,33"
	// }
	ReadObjectsInDB(from_tables string, params ...map[string]string) ([]map[string]string, error)
	//params: callback fun ej: fun([]map[string]string,error)
	// "ORDER_BY": "patient_name", "SORT":"DESC" DEFAULT "ASC"
	ReadStringDataAsyncInDB(r ReadDBParams, callback func([]map[string]string, error))
	ReadAnyDataAsyncInDB(r ReadDBParams, callback func([]map[string]interface{}, error))

	UpdateObjectsInDB(table_name string, data ...map[string]string) error
	DeleteObjectsInDB(table_name string, data ...map[string]string) error

	CreateTablesInDB(objects []*Object, result func(error))

	BackupDataBase(callback func(error))
}

// acción a ejecutar posteriormente
type Subsequently interface {
	ActionExecutedLater()
}
