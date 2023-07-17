package model

// FolderPath() string ej: func (Object) FolderPath() string {
// _, filename, _, _ := runtime.Caller(0)
// dir := filepath.Dir(filename)
// return filepath.ToSlash(dir)
// }
type Path interface {
	FolderPath() string
}
