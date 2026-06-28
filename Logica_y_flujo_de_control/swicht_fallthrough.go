package main

import "fmt"

func main() {
    // Leer entrada
    var alertLevel string
    fmt.Scanln(&alertLevel)
    
    // TODO: Escribe tu código a continuación
    // Crear una declaración switch con fallthrough para el sistema de alerta
    switch alertLevel {
	case "critical":
			fmt.Println("CRITICAL: System shutdown imminent!")
			fallthrough
		case "high": 
			fmt.Println("HIGH: Immediate attention required!")
			fallthrough
		case "medium":
			fmt.Println("MEDIUM: Please review when possible")
		case "low":
			fmt.Println("LOW: Informational only")
	}
}
