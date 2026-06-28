package main

import "fmt"

// Estructura Person con campos de nombre y edad
type Person struct {
	name string
	age  int
}

func main() {
	// Crear una estructura Person
	person := Person{
		name: "John",
		age:  25,
	}

	// Crear un puntero a la estructura Person
	personPtr := &person

	// Imprimir la información original de la persona
	fmt.Println("Original person:")
	fmt.Printf("Name: %s, Age: %d\n", person.name, person.age)

	// TODO: Usar el puntero para cambiar el nombre de la persona a "Jane" y la edad a 30
	personPtr.name = "Jane"
	// Sugerencia: Puedes acceder a los campos de la estructura a través de un puntero usando la notación de punto
	

	// Imprimir la información actualizada de la persona
	fmt.Println("Updated person:")
	fmt.Printf("Name: %s, Age: %d\n", person.name, person.age)
}