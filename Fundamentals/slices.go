package main

import "fmt"

func main() {
	// Aquí está nuestra porción original de frutas
	fruits := []string{"apple", "banana", "orange", "grape", "kiwi"}
	
	// TODO: Crear una nueva porción llamada 'firstThree' que contenga solo las primeras tres frutas
	firstThree := fruits[:3]
	// TODO: Crear una nueva porción llamada 'lastTwo' que contenga solo las últimas dos frutas
	lastTwo := fruits[len(fruits)-2:]
	// Imprimir los resultados
	fmt.Println("First three fruits:", firstThree)
	fmt.Println("Last two fruits:", lastTwo)
}
