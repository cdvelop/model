package model

type BackendHeaderHandler interface {
	HeaderBackendRequest() map[string]string
}

type FrontendHeaderHandler interface {
	GetFrontendHeaderValues() []string
	RunActionValueHeader(actions ...string) error
}
