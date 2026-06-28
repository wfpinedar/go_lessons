package main

import (
	"errors"
	"fmt"
)

// celsiusToFahrenheit convierte Celsius a Fahrenheit
// Devuelve un error si la temperatura está por debajo del cero absoluto (-273.15°C)
func celsiusToFahrenheit(celsius float64) (float64, error) {
	// TODO: Implementar la lógica de conversión
	// 1. Verificar si la temperatura está por debajo del cero absoluto (-273.15°C)
	if celsius < -273 {
		return 0, errors.New("Error: temperature below absolute zero")
	} else {
		// 2. Si es válida, convertir usando la fórmula: F = C × 9/5 + 32
		fahrenheit := (celsius * 9.0/5.0) + 32.0
		// 3. Devolver el valor y el error apropiados
		return fahrenheit, nil
	}

}

func main() {
	// Probar temperatura válida
	fmt.Println("Converting 25°C to Fahrenheit:")
	fahrenheit, err := celsiusToFahrenheit(25)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("25°C = %.2f°F\n", fahrenheit)
	}

	// Probar otra temperatura válida
	fmt.Println("\nConverting 0°C to Fahrenheit:")
	fahrenheit, err = celsiusToFahrenheit(0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("0°C = %.2f°F\n", fahrenheit)
	}

	// Probar temperatura inválida (por debajo del cero absoluto)
	fmt.Println("\nConverting -300°C to Fahrenheit:")
	fahrenheit, err = celsiusToFahrenheit(-300)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("-300°C = %.2f°F\n", fahrenheit)
	}
}