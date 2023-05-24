package model

type Request struct {
	//usuario que ejecuta la solicitud
	*User
	// paquetes de solicitud y respuesta
	Packages []Response
}
