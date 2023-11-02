package model

type FileAdapter interface {
	ReadFile(url_or_path string) ([]byte, error)
	DeleteFile(url_or_path string) error
}
