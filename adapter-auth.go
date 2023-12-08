package model

type AuthFrontendAdapter interface {
	FrontendCheckUser(result func(u *User, err string))
	AuthHandlerAdapter
}

type AuthBackendAdapter interface {
	// params ej: r *http
	BackendCheckUser(r_http any) (u *User, err string)
	AuthHandlerAdapter
}

type AuthHandlerAdapter interface {
	NameOfAuthHandler() string
	UserAuthNumber
}

type UserAuthNumber interface {
	// ej: 1 or 2 or 34 or 400.....
	UserAuthNumber() (number string, err string)
}
