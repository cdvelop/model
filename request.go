package model

type Request struct {
	//usuario que ejecuta la solicitud
	*User
	// paquetes de solicitud y respuesta
	Packages []Package
}

type Package struct {
	SkipMeInResponse bool     `json:"-"` //  saltarme al difundir
	Recipients       []string `json:"-"` // lista de ip, token o ids según app para el reenvió al finalizar solicitud
	TargetArea       bool     `json:"-"` // si es falso el destino por defecto es publico de lo contrario sera para todos quines tengan la misma area del usuario solicitante
	Response
}
