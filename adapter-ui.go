package model

type UI interface {
	UserInterface(options ...string) string
}

type FrontendHandler struct {
	AfterCreate
	AfterUpdate
	AfterDelete
}

type AfterCreate interface {
	SetObjectInDomAfterCreate(data ...map[string]string) (container_id, tags string)
}

type AfterUpdate interface {
	SetObjectInDomAfterUpdate(data ...map[string]string) (container_id, tags string)
}

type AfterDelete interface {
	SetObjectInDomAfterDelete(data ...map[string]string) (container_id, tags string)
}

// container_id ej: "contenedor-objeto"
// tags ej: <div>objeto...</div>
