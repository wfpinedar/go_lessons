package main

import "fmt"

func main() {
	// Mapa de artículos de inventario y sus cantidades
	inventory := map[string]int{
		"pen":    15,
		"pencil": 10,
		"paper":  25,
		"eraser": 5,
	}

	fmt.Println("Initial inventory:", inventory)

	// TODO: Eliminar la entrada 'pencil' del mapa de inventario
	delete(inventory, "pencil")
	

	fmt.Println("Updated inventory:", inventory)

	// Comprobar si 'pencil' todavía existe en el mapa
	_, exists := inventory["pencil"]
	fmt.Println("Does 'pencil' exist in inventory?", exists)
}
