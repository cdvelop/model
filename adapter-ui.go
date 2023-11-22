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

	ViewAdapter

	ResetViewObjectAdapter
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

type ViewAdapter interface {
	NameViewAdapter() string
	ContainerViewAdapter
	ItemViewAdapter
	ResetViewAdapter
}

// todo el contenido html por defecto del objeto
type ContainerViewAdapter interface {
	BuildContainerView(id, field_name string, allow_skip_completed bool) string
}

type ItemViewAdapter interface {
	BuildItemsView(all_data ...map[string]string) (html string)
}

// container_id ej: "contenedor-objeto"
// tags ej: <div>objeto...</div>
