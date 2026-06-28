package main

import "fmt"

func main() {
	// Esta es una cuadrícula de 3x3 representada como un array 2D
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	
	// TODO: Acceder e imprimir el valor en la fila 1, columna 2 (debería ser 6)
	// Recordatorio: los índices de los arrays comienzan en 0, por lo que la fila 1 es la segunda fila
	
	// Escribe tu código debajo de esta línea
	valor := grid[1][2]
    fmt.Println(valor)
}