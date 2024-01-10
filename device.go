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
	GamepadClientNotifyRegister(gamepadConfig any)
}
