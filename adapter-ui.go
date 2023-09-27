package model

type UI interface {
	UserInterface(options ...string) string
}

type FrontendHandler struct {
	AfterCreate
	AfterRead
	AfterUpdate
	AfterDelete
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

// container_id ej: "contenedor-objeto"
// tags ej: <div>objeto...</div>
