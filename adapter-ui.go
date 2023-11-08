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

	ViewReset
}

// type StoreData interface {
// 	GenerateStoreData(items ...any) (store_data map[string]interface{}, err error)
// }

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
	ViewHandlerName() string
	ContainerView
	ItemView
}

type ViewReset interface {
	ResetView()
}

// todo el contenido html por defecto del objeto
type ContainerView interface {
	BuildContainerView(id, field_name string, allow_skip_completed bool) string
}

type ItemView interface {
	BuildItemView(all_data ...map[string]string) (html string)
}

// container_id ej: "contenedor-objeto"
// tags ej: <div>objeto...</div>
