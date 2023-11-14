package model

type FileApi interface {
	// FileAddUploadApi(h *Handlers, source *Object, file_settings any) error
	//backend file_request ej: r *http.Request, w http.ResponseWriter
	//frontend file_request ej: blob file js.Value
	FileUpload(object_name, area_file string, file_request ...any) (out []map[string]string, err error)
	//params ej: id:123
	FilePath(params map[string]string) (file_path, area string, err error)
}

type FileDiskRW interface {
	FileGet(path string) (any, error)

	FileDelete(path string) error
}
