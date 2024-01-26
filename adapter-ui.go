package model

type FrontendHandler struct {
	AfterCreate
	AfterRead
	AfterUpdate
	AfterDelete

	AfterClickNotify

	NotifyBootData
	NotifyFileDelete

	ViewHandlerObject

	ResetFrontendObjectStateAdapter

	FormNotify
}

type AfterCreate interface {
	SetObjectInDomAfterCreate(data ...map[string]string) (err string)
}

type AfterRead interface {
	SetObjectInDomAfterRead(data ...map[string]string) (err string)
}

type AfterUpdate interface {
	SetObjectInDomAfterUpdate(data ...map[string]string) (err string)
}

type AfterDelete interface {
	SetObjectInDomAfterDelete(data ...map[string]string) (err string)
}

type AfterClickNotify interface {
	// data del objeto que fue cliqueado por el usuario
	NotifyUserHasClickedObject()
}

type NotifyFileDelete interface {
	NotifyFileDeletion()
}

type NotifyBootData interface {
	NotifyBootDataIsLoaded()
}

// todo el contenido html por defecto del objeto
type ContainerViewAdapter interface {
	BuildContainerView(id, field_name string, allow_skip_completed bool) string
}

type ItemViewAdapter interface {
	BuildItemsView(all_data ...map[string]string) (html string)
}

type FormNotify interface {
	NotifyFormIsOK()
	NotifyFormERR(err string)
}
