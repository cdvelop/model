package model

type FileHandler interface {
	// PrepareFileData(header_name, extension string, form_data map[string]string) map[string]string
	//header_name ej: fileHeader.Filename,user_area_file ej:1,5 extension ej: .jpeg
	RegisterNewFile(new *FileNewToStore, form_data map[string]string) (map[string]string, error)
	ConfigFile() *FileConfig
	UploadFolderPath(form_data map[string]string) string // carpeta de destino
	GenerateFileNameOnDisk() string                      // retornar nombre de archivo sin extension
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

	ImagenWidth  string // ej: 800
	ImagenHeight string // ej: 600

	RootFolder string //ej: static_files default "app_files"
	FileType   string // ej: imagen,video,document,pdf

	SavedAsBlobInDb bool // si se guardara en el motor de la base de datos, por defecto en el disco como archivo

	//field
	FieldNameWithObjectID string //ej: id_medicalhistory
	Name                  string //ej: endoscopia, pictures
	Legend                string //ej: Imágenes,Boletas etc
	DefaultEnableInput    bool   // si se necesita habilitado resetear el campo por defecto falso

}

type FileNewToStore struct {
	DescriptionInputName string //name in the input. whit extension ej: perrito.jpg,boleta_123.png
	UploadFolder         string //ej: ./files
	FileNameOnDisk       string // no extension ej: 1223, archivo_235, 001_file
	FileArea             string // area que tendrá el archivo ej: s,a,1,3,4
	Extension            string // ej .jpg,.mp4
	BlobData             []byte // data del archivo
}
