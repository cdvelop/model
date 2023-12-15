package model

type AuthFrontendAdapter interface {
	FrontendCheckUser(result func(u *User, err string))
	HeaderAuthenticationAdapter
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

type User struct {
	Token          string // token de sesión solicitante
	Number         string // numero correlativo de usuario
	Id             string // id usuario
	Ip             string
	Name           string
	Area           string // valor de carácter ej: a,s,p... key "" sin area
	AreaName       string // nombre del area
	AccessLevel    string // aquí valor numérico 0 a 255
	LastConnection string //time.Time

	R any //r *http.Request.
	W any //w http.ResponseWriter
}

type HeaderAuthenticationAdapter interface {
	AddHeaderAuthentication() Header
}

// default 2 minutes
type CookieExpiration struct {
	OneHour bool
	OneDay  bool
	Forever bool
}
