package model

type err struct {
	message string
}

func Error(texts ...string) *err {

	var message string
	for i, text := range texts {
		if i == 0 {
			message += text
		} else {
			message += " " + text
		}
	}

	return &err{message: message}
}

func (e err) Error() string {
	return e.message
}
