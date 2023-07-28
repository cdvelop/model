package model

type DataBaseAdapter interface {
	CreateObjectsInDB(table_name string, data ...map[string]string) error
	ReadObjectsInDB(table_name string, data ...map[string]string) ([]map[string]string, error)
	UpdateObjectsInDB(table_name string, data ...map[string]string) ([]map[string]string, error)
	DeleteObjectsInDB(table_name string, data ...map[string]string) ([]map[string]string, error)

	CreateTablesInDB(...*Object) error
}
