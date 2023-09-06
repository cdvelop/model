package model

type User struct {
	Token          string // token de sesión solicitante
	Id             string // id usuario
	Ip             string
	Name           string
	Area           string //0 sin area.. un que byte y uint8 son lo mismo en go la idea que aca sea un valor de carácter ej: a,s,p...
	AccessLevel    string // aquí valor numérico 0 a 255
	LastConnection string //time.Time
	// Packages       chan []Response
}
