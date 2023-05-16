package model

import (
	"time"
)

type User struct {
	Token          string // token de sesión solicitante
	Ip             string
	Name           string
	Area           byte  //0 sin area.. un que byte y uint8 son lo mismo en go la idea que aca sea un valor de carácter ej: a,s,p...
	AccessLevel    uint8 // aquí valor numérico 0 a 255
	Packages       chan []*Response
	LastConnection time.Time
}
