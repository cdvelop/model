package model

type TimeAdapter interface {
	// layout ej: 2006-01-02
	ToDay(layout string) string
}
