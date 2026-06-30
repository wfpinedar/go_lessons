package main

import (
	"fmt"
	"strconv"
)

// Definir la interfaz Shape, las estructuras Circle
type Shape interface {
	Area() float64
}

type Circle struct {
	Radio float64
}

// Implementar métodos Area() para ambas estructuras

type Rectangle struct {
	Width  float64
	Height float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radio * c.Radio
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Crear la función printShapeInfo

func printShapeInfo(s Shape) {
	fmt.Printf("Area: %g", s.Area())
}

func main() {
	// Leer entrada
	var shapeType string
	var dimension1 string
	var dimension2 string
	fmt.Scanln(&shapeType)
	fmt.Scanln(&dimension1)
	fmt.Scanln(&dimension2)

	// Convertir entradas de cadena a float64
	dim1, _ := strconv.ParseFloat(dimension1, 64)
	dim2, _ := strconv.ParseFloat(dimension2, 64)

	// TODO: Escribe tu código a continuación
	// Crear una instancia de forma apropiada basada en shapeType

	switch shapeType {
	case "circle":
		shape := Circle{dim1}
		printShapeInfo(shape)
	case "rectangle":
		shape := Rectangle{dim1, dim2}
		printShapeInfo(shape)

	}
	// Llamar a printShapeInfo con la forma creada

}
