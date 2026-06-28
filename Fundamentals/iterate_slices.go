package main

import "fmt"

func main() {
	// Un slice de frutas
	fruits := []string{"Apple", "Banana", "Cherry", "Dragon fruit", "Elderberry"}
	
	// TODO: Completa el bucle for usando range para iterar a través del slice de frutas
	// Imprime cada fruta con su posición como: "1. Apple"
	
	for index, value := range fruits {
		// Agrega tu código aquí para imprimir cada fruta con su posición
		fmt.Printf("%d. %s\n", index+1, value)
	}
}
