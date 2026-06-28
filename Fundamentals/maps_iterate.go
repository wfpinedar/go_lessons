package main

import (
	"fmt"
	"sort"
)

func main() {
	// Un mapa de calificaciones de estudiantes
	studentGrades := map[string]string{
		"Alice":  "A",
		"Bob":    "B",
		"Charlie": "B+",
		"David":  "A-",
		"Eva":    "C",
	}

	// Recopilar los nombres de los estudiantes (claves) en un slice para ordenarlos
	var names []string
	for name := range studentGrades {
		names = append(names, name)
	}
	sort.Strings(names)


	// TODO: Iterar a través de la lista ordenada de nombres de estudiantes
	// e imprimir el nombre y la calificación de cada estudiante en el formato:
	// "Estudiante: [nombre], Calificación: [calificación]"
    for _, name := range names {
        fmt.Printf("Student: %s, Grade: %s\n", name, studentGrades[name])
    }
}
