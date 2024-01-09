package model

type DevicePeripherals struct {
	KeyboardClientAdapter
	GamepadClientAdapter
}

type KeyboardClientAdapter interface {
	KeyboardClientDisable(disable bool)
}

type GamepadClientAdapter interface {
	//config button number and call function
	GamepadCallFunRegisterButton(*GamepadConfig)
}

type GamepadConfig struct {
	Connected      func()         //  al conectar gamepad
	Disconnected   func()         //  al desconectar gamepad
	ButtonAny      func()         //  al presionar cualquier botón
	ButtonSpecific map[int]func() // función a ejecutar según numero de botón especifico
}
