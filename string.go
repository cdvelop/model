package model

// Contains verifica si la cadena 'search' está presente en 'text'
func contains(text, search string) int {
	// Si la cadena de búsqueda es una cadena vacía, no puede haber coincidencias
	if search == "" {
		return 0
	}

	// Obtén la longitud de la cadena de búsqueda
	searchLen := len(search)

	// Inicializa el contador de coincidencias
	count := 0

	// Recorre el texto y cuenta la cantidad de coincidencias
	for i := 0; i <= len(text)-searchLen; i++ {
		if text[i:i+searchLen] == search {
			count++
		}
	}

	// Devuelve la cantidad de coincidencias encontradas
	return count
}
