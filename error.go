package model

type MyError struct {
	Message string
}

func (e MyError) Error() string {
	return e.Message
}
