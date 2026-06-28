package main

import "fmt"

func main() {
	// Slice de origen con nombres de frutas
	source := []string{"apple", "banana", "cherry", "date"}
	
	// Crear un slice de destino con capacidad para 3 elementos
	destination := make([]string, 3)
	
	// TODO: Usar la función copy para copiar elementos del origen al destino
	copy(destination, source)
	
	// Imprimir el slice de destino
	fmt.Println("Destination slice:", destination)
}
