package main

import "fmt"

func main() {
	// TODO: Crea un literal de mapa llamado 'capitals' que mapee países a sus ciudades capitales
	// Incluye al menos estos tres pares:
	// "France" -> "Paris"
	// "Japan" -> "Tokyo"
	// "Brazil" -> "Brasilia"
	capitals := map[string]string{
        "France": "Paris",
	    "Japan": "Tokyo",
	    "Brazil": "Brasilia",
    }
	// Esto imprimirá el mapa
	fmt.Println(capitals)
	
	// Esto imprimirá la capital de Japón
	fmt.Println("The capital of Japan is:", capitals["Japan"])

	scores := make(map[string]int, 10)
	fmt.Println(scores)
}
