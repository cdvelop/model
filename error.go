package model

import "strconv"

type err struct {
	message string
}

//only support: error, int, int64, string, map[string]interface{}
func Error(messages ...any) *err {

	var text string
	var space string
	for i, msg := range messages {
		var txt string

		switch v := msg.(type) {
		case error:
			txt = v.Error()

		case int:
			txt = strconv.Itoa(v)

		case int64:
			txt = strconv.FormatInt(v, 10)

		case string:
			txt = v

		case map[string]interface{}:
			var comma string
			for t := range v {
				txt += comma + t
				comma = ","
			}

		}

		if txt == "" {
			txt = "¡valor n°:" + strconv.Itoa(i) + " no soportado en mensaje error!"
		}

		text += space + txt
		space = " "
	}

	return &err{message: text}
}

func (e err) Error() string {
	return e.message
}
