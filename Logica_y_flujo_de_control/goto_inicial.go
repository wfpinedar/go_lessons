package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Leer entrada
    var maxRetriesStr string
    var successAttemptStr string
    fmt.Scanln(&maxRetriesStr)
    fmt.Scanln(&successAttemptStr)
    
    // Analizar entradas a enteros
    maxRetries, _ := strconv.Atoi(maxRetriesStr)
    successAttempt, _ := strconv.Atoi(successAttemptStr)
    
    // Inicializar contador de intentos
    attempt := 1
    
    // TODO: Escribe tu código a continuación
    // Usar la declaración goto con una etiqueta para implementar la lógica de reintento
		reatemp:
		if successAttempt == attempt {
			fmt.Printf("Attempt %d succeeded\n", attempt) 
		} else if attempt > maxRetries {
			fmt.Println("All attempts failed")
		} else {
			fmt.Printf("Attempt %d failed\n", attempt)
			attempt = attempt + 1
			goto reatemp
		}

}

