package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    // Leer entrada
    var dimensions string
    var targetStr string
    fmt.Scanln(&dimensions)
    fmt.Scanln(&targetStr)
    
    // Analizar dimensiones
    dimParts := strings.Split(dimensions, ",")
    rows, _ := strconv.Atoi(dimParts[0])
    cols, _ := strconv.Atoi(dimParts[1])
    
    // Analizar número objetivo
    target, _ := strconv.Atoi(targetStr)
    
    // Cuadrícula predefinida
    grid := [][]int{
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10, 11, 12},
    }
    
    // TODO: Escribe tu código a continuación
    // Usar bucles anidados con ruptura etiquetada para buscar el objetivo
	found := false
    scape:
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == target {
					fmt.Printf("Found %d at position (%d, %d)", target, row, col)
					found = true
					break scape
				}

		}
	}
	if !found {
		fmt.Printf("Target %d not found", target)
	}
    
}