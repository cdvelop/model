package model

type IdHandler interface {
	GetNewID() (id, err string)
}

type ReadParams struct {
	FROM_TABLE      string   //ej: "users,products" or: public.reservation, public.patient"
	ID              string   // unique search
	SELECT          string   // "name, phone, address" default *
	WHERE           []string //"patient.id_patient = reservation.id_patient AND reservation.id_staff = '2'"
	SEARCH_ARGUMENT string   //"1,4,33"
	ORDER_BY        string   // name,phone,address
	SORT_DESC       bool     //default ASC
	LIMIT           int      // 10, 5, 100. note: Postgres y MySQL: "LIMIT 10", SQLite: "LIMIT 10 OFFSET 0" OR "" no limit
	RETURN_ANY      bool     // default string return []map[string]string, any = []map[string]interface{}
}

type ReadResult struct {
	DataString []map[string]string
	DataAny    []map[string]any
	Error      string
}

type DataBaseAdapter interface {
	IdHandler
	RunOnClientDB() bool //verdadero corren en el cliente ej browser. por defecto falso corre en el servidor
	// items support in server db: []map[string]string, map[string]string
	CreateObjectsInDB(table_name string, backup_required bool, items any) (err string)

	ReadSyncDataDB(p ReadParams, data ...map[string]string) (result []map[string]string, err string)
	ReadAsyncDataDB(p ReadParams, callback func(r ReadResult))

	UpdateObjectsInDB(table_name string, data ...map[string]string) (err string)
	DeleteObjectsInDB(table_name string, data ...map[string]string) (err string)

	CreateTablesInDB(objects []*Object, result func(err string))

	ClearAllTableDataInDB(tables ...string) (err string)
	BackupDataBase(callback func(err string))
}
