package model

type UI interface {
	UserInterface(options ...string) string
}

type FrontendHandler struct {
	AfterCreate
	AfterRead
	AfterUpdate
	AfterDelete

	AfterClicked

	NotifyBootData

	ViewHandler
}

type AfterCreate interface {
	SetObjectInDomAfterCreate(data ...map[string]string) error
}

type AfterRead interface {
	SetObjectInDomAfterRead(data ...map[string]string) error
}

type AfterUpdate interface {
	SetObjectInDomAfterUpdate(data ...map[string]string) error
}

type AfterDelete interface {
	SetObjectInDomAfterDelete(data ...map[string]string) error
}

type AfterClicked interface {
	// data del objeto que fue cliqueado por el usuario
	UserClicked(data map[string]string)
}

type NotifyBootData interface {
	NotifyBootDataIsLoaded()
}

type ViewHandler interface {
	ViewComponentName() string
	BuildTag() string
	HtmlContainer
}

// todo el contenido html por defecto del objeto
type HtmlContainer interface {
	HtmlContainer() string
}

// container_id ej: "contenedor-objeto"
// tags ej: <div>objeto...</div>
