package model

type DevicePeripherals struct {
	KeyboardClientAdapter
	GamepadClientAdapter
}

type KeyboardClientAdapter interface {
	KeyboardClientDisable(disable bool)
}

type GamepadClientAdapter interface {
	GamepadRegister(notifyConnection, pressAnyButton, pressSpecificButtons any)
}
