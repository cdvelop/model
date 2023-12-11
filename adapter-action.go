package model

type BeforeAction interface {
	ExecuteBeforeAction(h *Handlers)
}
