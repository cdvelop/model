package model

type BackendAuthHandler interface {
	AddBootDataActions() BootActions
}

type FrontendAuthHandler interface {
	ContainerIdBootData() string // ej:boot-data <div id="boot-data"></div>
}

type BootActions struct {
	DataBootActions string
}

type DefaultAuthHandler struct{}

func (DefaultAuthHandler) AddBootDataActions() BootActions {
	return BootActions{
		DataBootActions: "",
	}
}

func (DefaultAuthHandler) ContainerIdBootData() string {
	return ""
}
