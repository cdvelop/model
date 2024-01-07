package model

type BackupHandlerAdapter interface {
	//backup_type ej: "create","delete","update"
	BackupOneObjectType(backup_type, table_name string, data any)
}
