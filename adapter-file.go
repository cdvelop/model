package model

type BackendHandler struct {
	BootResponse

	CreateApi
	ReadApi
	UpdateApi
	DeleteApi

	FileHandler
}

type BootResponse interface {
	AddBootResponse(u *User) ([]Response, error)
}

type CreateApi interface {
	Create(u *User, data ...map[string]string) error
}

type ReadApi interface {
	Read(u *User, data ...map[string]string) ([]map[string]string, error)
}

type UpdateApi interface {
	Update(u *User, data ...map[string]string) ([]map[string]string, error)
}

type DeleteApi interface {
	Delete(u *User, data ...map[string]string) ([]map[string]string, error)
}

type FileHandler interface {
	// PrepareFileData(header_name, extension string, form_data map[string]string) map[string]string
	//header_name ej: fileHeader.Filename, extension ej: .jpeg
	RegisterNewFile(header_name, upload_folder, file_name, extension string, form_data map[string]string) (map[string]string, error)
	ConfigFile() *FileConfig
	UploadFolderPath(form_data map[string]string) string // carpeta de destino
	FileName() string                                    // retornar nombre de archivo sin extension
	FileApi
}

type FileApi interface {
	// ej: id:123
	GetFilePathByID(params map[string]string) (file_path, area string, err error)
}

type FileConfig struct {
	MaximumFilesAllowed int64  // numero máximo de archivos permitidos ej: 4,100 etc
	InputNameWithFiles  string // nombre del campo con los archivos multipart del formulario ej: files
	MaximumFileSize     int64  // tamaño máximo de todos los archivos
	MaximumKbSize       int64  // tamaño máximo individual kb ej: 100
	AllowedExtensions   string // exenciones permitidas ej: ".jpg, .png, .jpeg"

	RootFolder string //ej: static_files default "app_files"
	FileType   string // ej: imagen,video,document,pdf

	//field
	IdFieldName    string //ej: id_medicalhistory
	Name           string //ej: endoscopia, pictures
	Legend         string //ej: Imágenes,Boletas etc
	TabIndexNumber string // ej 3,5.. posición con respecto a otros campos del formulario para uso en tabindex

}
