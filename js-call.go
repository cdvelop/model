package model

type CallJsFunWithParameters struct {
	FuncNameCall string         // función a llamar
	Enable       bool           // estado por defecto a enviar
	AddParams    map[string]any // parámetros adicionales

	all_params map[string]any // uso interno
}

func (c *CallJsFunWithParameters) ExecuteJsFun(o *Object) (err string) {
	// Crear el mapa que contendrá todos los parámetros
	c.all_params = make(map[string]interface{})

	// Añadir los parámetros específicos de la función
	c.all_params["enable"] = c.Enable
	c.all_params["query"] = o.QuerySelectorObject(o.ModuleName, o.ObjectName)

	// Añadir los parámetros adicionales proporcionados por el usuario
	for key, value := range c.AddParams {
		c.all_params[key] = value
	}

	// Llamar a la función JavaScript
	return o.CallFunction(c.FuncNameCall, c.all_params)

}
