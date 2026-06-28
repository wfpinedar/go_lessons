package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Leer entrada
	var rowsInput string
	var skipConditionInput string
	fmt.Scanln(&rowsInput)
	fmt.Scanln(&skipConditionInput)

	// Analizar entradas
	numRows, _ := strconv.Atoi(rowsInput)
	skipCondition, _ := strconv.Atoi(skipConditionInput)

	// Matriz de datos 2D proporcionada
	data := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	// TODO: Escribe tu código a continuación
	// Usar bucles anidados con 'continue' etiquetado para procesar los datos
	// Verificar cada fila en busca del número de condición de omisión
	// Imprimir mensajes apropiados según si se encuentra la condición
outer:
	for row := 0; row < numRows; row++ {
		mesagge := "Processing row " + strconv.Itoa(row) + ":"
		for col := 0; col < len(data[row]); col++ {
			if data[row][col] == skipCondition {
				fmt.Printf("Skipping row %d due to condition\n", row)
				continue outer
			} else {
				mesagge = mesagge + " " + strconv.Itoa(data[row][col])
			}
		}
		fmt.Printf("%s\n", mesagge)
	}
}
