package main

import "fmt"

func main() {
	// TODO: Crear un mapa llamado 'favoriteColors' usando make
	// Las claves deben ser cadenas (nombres) y los valores también deben ser cadenas (colores)
	favoriteColors := make(map[string]string)
	// Añadiendo algunos pares clave-valor al mapa
	favoriteColors["Alice"] = "Blue"
	favoriteColors["Bob"] = "Green"
	favoriteColors["Charlie"] = "Red"
	
	// Imprimir el mapa
	fmt.Println("Favorite colors:", favoriteColors)
	
	// Imprimir el color favorito de Bob
	fmt.Println("Bob's favorite color is", favoriteColors["Bob"])
}
