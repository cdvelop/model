package model

type ResetViewAdapter interface {
	ResetAdapterView() (err string)
}

type ResetViewObjectAdapter interface {
	ResetObjectView() (err string)
}
