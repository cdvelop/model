package model

type AuthAdapter interface {
	UserAuthNumber
	LoginUser
}

// ej: 1 or 2 or 34 or 400.....
type UserAuthNumber interface {
	UserAuthNumber() (number string, err string)
}

type LoginUser interface {
	// params ej: r *http
	GetLoginUser(params any) (u *User, err string)
}
