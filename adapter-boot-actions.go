package model

type FrontendBootActions interface {
	// ContainerIdBootData() string // ej:boot-data <div id="boot-data"></div>
}

type BootActions struct {
	JsonBootActions string
}

// type DefaultAuthHandler struct{}

// func (DefaultAuthHandler) AddBootResponse() []Response {
// 	return []Response{}
// }

// func (DefaultAuthHandler) ContainerIdBootData() string {
// 	return ""
// }
