package model

type DiskAdapter interface {
	ReadFileFrom(url_or_path string) ([]byte, error)
}
