package model

type CallJsOptions struct {
	NameJsFunc         string
	Enable             bool //parámetro estado por defecto a enviar
	NotSendQueryObject bool //no envía querySelector Objeto por defecto siempre se envía

	Params map[string]any // parámetros

	ResultInt    bool // resultado en formato int
	ResultString bool // resultado en string
}

// enviar con parámetros enable y query del objeto
func (c CallJsOptions) CallWithEnableAndQueryParams(o *Object) (result any, err string) {

	// Añadir los parámetros específicos de la función
	c.Params["enable"] = c.Enable

	if !c.NotSendQueryObject {
		c.Params["query"] = o.QuerySelectorObject(o.ModuleName, o.ObjectName)
	}

	return o.CallFunction(c.NameJsFunc, c)

}
