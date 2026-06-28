package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Leer entrada
	var threatLevel string
	var maxZonesStr string
	fmt.Scanln(&threatLevel)
	fmt.Scanln(&maxZonesStr)

	// Convertir el número máximo de zonas a entero
	maxZones, _ := strconv.Atoi(maxZonesStr)

	// Datos de las zonas de seguridad
	zones := [][]string{
		{"low", "medium", "low"},
		{"medium", "high", "low"},
		{"critical", "medium", "high"},
		{"low", "critical", "medium"},
	}

	// TODO: Escribe tu código a continuación
	// Usa bucles anidados con 'labeled break' para buscar la amenaza
	// Imprime la ubicación de la amenaza si se encuentra, o "Amenaza no encontrada en las zonas buscadas" si no se encuentra
	// Luego usa switch con 'fallthrough' para las alertas de seguridad
	found := false
zonacatch:
	for zone := 0; zone < maxZones; zone++ {
		for pos := 0; pos < len(zones[zone]); pos++ {
			if threatLevel == zones[zone][pos] {
				found = true
				fmt.Printf("Threat found at zone %d position %d\n", zone, pos)
				break zonacatch
			}
		}
	}
	if !found {
		fmt.Println("Threat not found in searched zones")
	}

	switch threatLevel {
	case "critical":
		fmt.Println("CRITICAL: Lockdown initiated!")
		fallthrough
	case "high":
		fmt.Println("HIGH: Security breach detected!")
		fallthrough
	case "medium":
		fmt.Println("MEDIUM: Increased monitoring active")
	case "low":
		fmt.Println("LOW: Standard security protocols")
	}
}
