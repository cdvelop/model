package model

type Response struct {
	Action  string              `json:"a,omitempty"` //acción a realizar con la solicitud: create, read, update, delete o error
	Object  string              `json:"o,omitempty"` //nombre de la tabla u objeto controlador hacia donde va la solicitud
	Message string              `json:"g,omitempty"` //mensaje para el usuario de la solicitud
	Data    []map[string]string `json:"d,omitempty"` //data entrada y respuesta json
}

type CutResponse struct {
	//Action,Object,Module,Message
	//ej: ["create","user","ModuleUsers","ok"]
	CutOptions []string  `json:"o"`
	CutData    []CutData `json:"d"`
}

type CutData struct {

	// ej Modelo Objeto: usuario := Object{Name: "Usuario",Fields: []Field{
	// 		{Name: "name"}, //0
	// 		{Name: "email"},//1
	// 		{Name: "phone"},//2
	// 	},}
	// ej data normal con todos los campos: {"name":"John Doe","email":"johndoe@example.com","phone":"555"}
	// version recortada Data: {"John Doe","johndoe@example.com","555"}
	// Index Objeto al codificar = {"0:0","1:1","2:2"}
	// ej no mail: {"marcel", "777"}
	// Index al codificar = {"0:0","1:2"}
	Index map[uint8]uint8 `json:"i"`
	//Data ej en mapa: "Data":[{"id":"222","name":"manzana","valor":"1200"}]
	// ahora: "Data":["222","manzana","1200"]
	Data []string `json:"d"`
}
