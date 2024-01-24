package model

type ViewHandlerObject interface {
	ViewHandlerName() string
	ContainerViewAdapter
	ItemViewAdapter
	ResetViewHandlerObjectAdapter
	NotifyStatusChangeAfterClicking()
}

func (o *Object) AddViewHandlerObject(viewHandlerObject any) {

	if vh, ok := viewHandlerObject.(ViewHandlerObject); ok {
		o.FrontHandler.ViewHandlerObject = vh
	}

}
