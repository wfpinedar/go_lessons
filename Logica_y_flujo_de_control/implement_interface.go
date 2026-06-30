package main

import (
	"fmt"
	"strconv"
)

// TODO: Escribe tu código a continuación
// 1. Definir la interfaz Vehicle
type Vehicle interface {
	GetInfo() string
}

// 2. Definir la estructura Car
type Car struct {
	Brand string
	Speed int
}

// 3. Implementar el método GetInfo() para Car
func (c Car) GetInfo() string {
	return "Brand: " + c.Brand + " Speed: " + strconv.Itoa(c.Speed) + " km/h"
}

func main() {
	// Leer entrada
	var brand string
	var speedStr string
	fmt.Scanln(&brand)
	fmt.Scanln(&speedStr)

	// Convertir la cadena de velocidad a entero
	speed, _ := strconv.Atoi(speedStr)

	// 4. Crear una instancia de Car y llamar a GetInfo()
	car := Car{brand, speed}
	result := car.GetInfo()
	// Imprimir el resultado
	fmt.Println(result)
}
