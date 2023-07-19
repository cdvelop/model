package model

type DataBase interface {
	CreateObjects(table_name string, all_data ...map[string]string) error
	ReadDbData(params map[string]string) ([]map[string]string, error)
	UpdateObjects(table_name string, all_data ...map[string]string) error
	DeleteObjects(table_name string, all_data ...map[string]string) error
}
