package model

type IdHandler interface {
	GetNewID() string
}

type DataBaseAdapter interface {
	IdHandler
	RunOnClientDB() bool //verdadero corren en el cliente ej browser. por defecto falso corre en el servidor
	// items support in server db: []map[string]string, map[string]string
	CreateObjectsInDB(table_name string, backup_required bool, items any) (err string)

	// from_tables ej: "users,products" or: public.reservation, public.patient"
	// params: map[string]string ej:{
	// LIMIT: 10, 5, 100. note: Postgres y MySQL: "LIMIT 10", SQLite: "LIMIT 10 OFFSET 0" OR "" no limit
	// ORDER_BY: name,phone,address
	// SELECT: "name, phone, address" default *
	// WHERE: "patient.id_patient = reservation.id_patient AND reservation.id_staff = '2'"
	// ARGS: "1,4,33"
	// }
	ReadObjectsInDB(from_tables string, params ...map[string]string) (result []map[string]string, err string)
	//params: callback fun ej: fun([]map[string]string,error)
	// "ORDER_BY": "patient_name", "SORT":"DESC" DEFAULT "ASC"
	ReadStringDataAsyncInDB(r ReadDBParams, callback func(result []map[string]string, err string))
	ReadAnyDataAsyncInDB(r ReadDBParams, callback func(result []map[string]interface{}, err string))

	UpdateObjectsInDB(table_name string, data ...map[string]string) (err string)
	DeleteObjectsInDB(table_name string, data ...map[string]string) (err string)

	CreateTablesInDB(objects []*Object, result func(err string))

	BackupDataBase(callback func(err string))
	ClearAllTableDataInDB(tables ...string) (err string)
}
