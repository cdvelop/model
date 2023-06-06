package model

type Icon struct {
	Id      string   //ej: icon-lupa, icon-btn-save, icon-btn-new
	ViewBox string   //ej: "0 0 16 16", "0 0 448 512"
	Paths   []string //ej: m11.51 11..,m7.051..,m3.267 15...
}

// ej: <symbol id="icon-btn-save" viewBox="0 0 16 16"><path fill="currentColor" d="M14 0h-2z" /></symbol>
func (i Icon) BuildSymbol() string {
	out := `<symbol id="` + i.Id + `" viewBox="` + i.ViewBox + `">`

	for _, d := range i.Paths {

		out += `<path fill="currentColor" d="` + d + `" />`

	}

	out += `</symbol>`
	return out
}
