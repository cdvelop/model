package model

type DevicePeripherals struct {
	KeyboardClientAdapter
}

type KeyboardClientAdapter interface {
	KeyboardClientDisable(disable bool)
}
