package main

import (
	"fmt"
	"strconv"
	"strings"
)

// processData acepta interface{} y puede manejar cualquier tipo
func processData(data interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", data, data)
}

func main() {
	// Leer entrada
	var dataType string
	var valueStr string
	fmt.Scanln(&dataType)
	fmt.Scanln(&valueStr)
	
	// Convertir valueStr al tipo apropiado basado en dataType
	switch dataType {
	case "int":
		// Convertir a entero
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			fmt.Println("Error al convertir a int:", err)
			return
		}
		processData(value)
		
	case "string":
		// Usar como cadena directamente
		processData(valueStr)
		
	case "bool":
		// Convertir a booleano
		value, err := strconv.ParseBool(valueStr)
		if err != nil {
			fmt.Println("Error al convertir a bool:", err)
			return
		}
		processData(value)
		
	case "slice":
		// Convertir valores separados por comas a slice de enteros
		parts := strings.Split(valueStr, ",")
		slice := make([]int, 0, len(parts))
		
		for _, part := range parts {
			num, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				fmt.Println("Error al convertir elemento del slice:", err)
				return
			}
			slice = append(slice, num)
		}
		processData(slice)
		
	default:
		fmt.Println("Tipo de dato no soportado:", dataType)
	}
}
