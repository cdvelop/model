package model

type FetchAdapter interface {
	//method ej: GET,POST
	// endpoint ej: http://localhost/upload
	// object ej: imagen
	// body_rq any ej: []map[string]string, map[string]string
	SendOneRequest(method, endpoint, object string, body_rq any, response func([]map[string]string, error))
	// only POST method
	SendAllRequests(endpoint string, body_rq []Response, response func([]Response, error))
}
