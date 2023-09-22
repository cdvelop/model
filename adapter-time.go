package model

type TimeAdapter interface {
	// ej: 2023-12-30
	PresentDate(layout string) string
}
