package model

type FileConfig struct {
	MaximumFilesAllowed int64  // numero máximo de archivos permitidos ej: 4,100 etc
	InputNameWithFiles  string // nombre del campo con los archivos multipart del formulario ej: files
	MaximumFileSize     int64  // tamaño máximo de todos los archivos
	MaximumKbSize       int64  // tamaño máximo individual kb ej: 100
	AllowedExtensions   string // exenciones permitidas ej: ".jpg, .png, .jpeg"

	RootFolder string //ej: static_files default "app_files"
	FileType   string // ej: imagen,video,document,pdf

	//field
	IdFieldName        string //ej: id_medicalhistory
	Name               string //ej: endoscopia, pictures
	Legend             string //ej: Imágenes,Boletas etc
	DefaultEnableInput bool   // si se necesita habilitado resetear el campo por defecto falso

}
