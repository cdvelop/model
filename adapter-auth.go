package model

type AuthAdapter interface {
	UserAuthNumber
}

// ej: 1 or 2 or 34 or 400.....
type UserAuthNumber interface {
	UserAuthNumber() (string, error)
}
