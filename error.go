package model

import (
	"strconv"
)

type err struct {
	message string
}

// only support: error, int, int64, string, map[string]interface{}
func Error(messages ...any) *err {

	// fmt.Printf("mensaje (%T): \n", messages)
	var text string
	var space string
	for i, msg := range messages {
		// fmt.Printf("mensaje %d (%T): %v\n", i, msg, msg)
		var txt string

		// Verificar si msg es una cadena vacía
		if str, ok := msg.(string); ok && str == "" {
			txt = "''"
		} else {

			switch v := msg.(type) {
			case error:
				txt = v.Error()

			case []string:
				var comma string
				var new string
				for _, x := range v {
					new += comma + x
					comma = ","
				}
				txt += new
				if new == "" {
					txt += "''"
				}

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
