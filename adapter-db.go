package model

type DataBaseAdapter interface {
	CreateObjectsInDB(table_name string, data ...map[string]string) error

	// from_tables ej: "users,products" or: public.reservation, public.patient"
	// data ... map[string]string ej:{
	// LIMIT: 10, 5, 100. note: Postgres y MySQL: "LIMIT 10", SQLite: "LIMIT 10 OFFSET 0" OR "" no limit
	// ORDER_BY: name,phone,address
	// SELECT: "name, phone, address" default *
	// WHERE: "patient.id_patient = reservation.id_patient AND reservation.id_staff = '2'"
	// ARGS: "1,4,33"
	// }
	ReadObjectsInDB(from_tables string, data ...map[string]string) ([]map[string]string, error)
	UpdateObjectsInDB(table_name string, data ...map[string]string) ([]map[string]string, error)
	DeleteObjectsInDB(table_name string, data ...map[string]string) ([]map[string]string, error)

	CreateTablesInDB(...*Object) error
}
