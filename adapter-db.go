package model

const PREFIX_ID_NAME = "id_"

type IdHandler interface {
	GetNewID() (id, err string)
}

type ReadParams struct {
	FROM_TABLE string //ej: "users,products" or: public.reservation, public.patient"
	SELECT     string // "name, phone, address" default *

	WHERE         []map[string]string //ej: "patient.id_patient=reservation.id_patient, (OR) reservation.id_staff = '2'"
	AND_CONDITION bool                // OR default se agrega AND si es true

	ID string // unique search (usado en indexdb)

	ORDER_BY  string // name,phone,address
	SORT_DESC bool   //default ASC
	LIMIT     int    // 10, 5, 100. note: Postgres y MySQL: "LIMIT 10", SQLite: "LIMIT 10 OFFSET 0" OR "" no limit

	RETURN_ANY bool // default string return []map[string]string, any = []map[string]interface{}
}

type ReadResults struct {
	ResultsString []map[string]string
	ResultsAny    []map[string]any
}

// backup_required ej si se esta en frontend se envía también al backend la solicitud
type DataBaseAdapter interface {
	IdHandler
	RunOnClientDB() bool //verdadero corren en el cliente ej browser. por defecto falso corre en el servidor
	// items support in server db: []map[string]string, map[string]string
	CreateObjectsInDB(table_name string, backup_required bool, items any) (err string)

	ReadSyncDataDB(p ReadParams, data ...map[string]string) (result []map[string]string, err string)
	ReadAsyncDataDB(p ReadParams, callback func(r *ReadResults, err string))

	UpdateObjectsInDB(table_name string, backup_required bool, data ...map[string]string) (err string)
	DeleteObjectsInDB(table_name string, backup_required bool, data ...map[string]string) (err string)

	CreateTablesInDB(objects []*Object, result func(err string))

	ClearAllTableDataInDB(tables ...string) (err string)
	// BackupDataBase(callback func(err string))
}
