package model

type Response struct {
	Type    string              `json:"t,omitempty"` //tipo solicitud respuesta: create, read, update, delete o error
	Object  string              `json:"o,omitempty"` //nombre de la tabla u objeto controlador hacia donde va la solicitud
	Module  string              `json:"m,omitempty"` //nombre del module controlador hacia donde va la solicitud
	Message string              `json:"g,omitempty"` //mensaje para el usuario de la solicitud
	Data    []map[string]string `json:"d,omitempty"` //data entrada y respuesta json

	SkipMeInResponse bool     `json:"-"` //  saltarme al difundir
	Recipients       []string `json:"-"` // lista de ip, token o ids según app para el reenvió al finalizar solicitud
	TargetArea       bool     `json:"-"` // si es falso el destino por defecto es publico de lo contrario sera para todos quines tengan la misma area del usuario solicitante
}

type CutResponse struct {
	//Type,Object,Module,Message
	//ej: ["create","user","ModuleUsers","ok"]
	CutOptions []string  `json:"o"`
	CutData    []CutData `json:"d"`
}
