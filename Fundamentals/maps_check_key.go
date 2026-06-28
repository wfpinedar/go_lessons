package main

import "fmt"

func main() {
	// Mapa de calificaciones de estudiantes
	grades := map[string]int{
		"John":  85,
		"Sarah": 92,
		"Mike":  78,
		"Lisa":  95,
	}

	// Estudiante a verificar
	studentToCheck := "Emma"

	// TODO: Verificar si studentToCheck existe en el mapa de calificaciones
	// y almacenar el resultado en una variable llamada 'exists'
	// Sugerencia: Usa el modismo de coma-ok


	// Imprimir el resultado
	if _, exists := grades[studentToCheck]; exists {
		fmt.Printf("%s's grade exists in the map\n", studentToCheck)
	} else {
		fmt.Printf("%s's grade does not exist in the map\n", studentToCheck)
	}
}