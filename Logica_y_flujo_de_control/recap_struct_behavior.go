package main

import (
	"fmt"
	"strconv"
)

// TODO: Define tu structura Rectangle aquí
type Rectangle struct {
    Width float64
	Height float64
}
// TODO: Implementa el método area con un receptor de valor aquí
func (r Rectangle) area() float64 {
	return r.Width * r.Height
}
// TODO: Implementa el método scale con un receptor de puntero aquí
func (r *Rectangle) scale(scaleFactor float64) {
	r.Width = r.Width * scaleFactor
	r.Height = r.Height * scaleFactor
}
func main() {
	// Leer entrada
	var widthStr string
	var heightStr string
	var scaleStr string
	fmt.Scanln(&widthStr)
	fmt.Scanln(&heightStr)
	fmt.Scanln(&scaleStr)
	
	// Analizar entradas a float64
	width, _ := strconv.ParseFloat(widthStr, 64)
	height, _ := strconv.ParseFloat(heightStr, 64)
	scaleFactor, _ := strconv.ParseFloat(scaleStr, 64)
	
	// TODO: Crea una instancia de Rectangle e implementa la solución a continuación
	rectangle := Rectangle{width, height}
	// Imprimir área inicial
	// fmt.Printf("Initial area: %.0f\n", initialArea) // Usa este formato para la salida
	initialArea := rectangle.area()
	fmt.Printf("Initial area: %.0f\n", initialArea)
	// Escalar el rectángulo
	rectangle.scale(scaleFactor)
	scaledArea := rectangle.area()
	// Imprimir área escalada
	// fmt.Printf("Scaled area: %.0f\n", scaledArea) // Usa este formato para la salida
	fmt.Printf("Scaled area: %.0f\n", scaledArea)
}