package model

type Module struct {
	//nombre modulo ej: chat,patient,user
	ModuleName string
	//Titulo que vera el usuario ej: "Modulo Fotografía"
	Title string
	// id icono para utilizar en sprite svg ej: icon-info
	IconID string

	//interfaz usuario modulo
	UI

	// ej:	`<div class="target-module">
	// 	<select name="select">
	// 		<option value="value1">Value 1</option>
	// 		<option value="value2" selected>Value 2</option>
	// 	</select>
	// </div>`
	HeaderInputTarget string

	//areas soportadas por el modulo ej: 'a','t','x'
	Areas []byte
	// objetos o componentes que contiene el modulo ej: patient,user,datalist,search....
	Objects []*Object

	// tipo de entradas usadas en el modulo
	Inputs []*Input
}
