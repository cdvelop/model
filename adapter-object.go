package model

type ObjectAdapter interface {
	GetObjectByName(name_to_search string) (*Object, error)
}
