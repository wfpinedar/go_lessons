package main

import "fmt"

func main() {
	// Un mapa de calificaciones de estudiantes
	studentGrades := map[string]string{
		"John":  "B+",
		"Emma":  "A-",
		"Sarah": "A",
		"Mike":  "C",
	}
	
	// TODO: Acceder y almacenar la calificación de Emma en una variable llamada emmaGrade
	grade := studentGrades["Emma"]
	// TODO: Imprimir la calificación de Emma con un mensaje como: "La calificación de Emma es: A-"
	fmt.Printf("Emma's grade is: %s", grade)
}
