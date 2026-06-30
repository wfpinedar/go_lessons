package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: Escribe tu código a continuación
// 1. Definir la interfaz Speaker
type Speaker interface {
	Speak()
}

// 2. Definir las estructuras Person y Parrot
type Person struct {
	Name string
}
type Parrot struct {
	Name string
}

// 3. Implementar los métodos Speak() para ambas estructuras
func (p Person) Speak() {
	fmt.Printf("Hello, I'm %s\n", p.Name)
}
func (p Parrot) Speak() {
	fmt.Printf("Squawk! %s says hello!\n", p.Name)
}

func makeAllSpeak(orator Speaker) {
	orator.Speak()
}

func main() {
	// Leer entrada
	var numSpeakersStr string
	var speakerTypesStr string
	var speakerNamesStr string

	fmt.Scanln(&numSpeakersStr)
	fmt.Scanln(&speakerTypesStr)
	fmt.Scanln(&speakerNamesStr)

	// Analizar entrada
	numSpeakers, _ := strconv.Atoi(numSpeakersStr)
	speakerTypes := strings.Split(speakerTypesStr, ",")
	speakerNames := strings.Split(speakerNamesStr, ",")

	// 4. Crear la función makeAllSpeak
	// 5. Crear oradores basados en la entrada y almacenarlos en un slice
	var orators []Speaker
	for i, tipo := range speakerTypes {
		switch tipo {
		case "person":
			orators = append(orators, Person{speakerNames[i]})
		case "parrot":
			orators = append(orators, Parrot{speakerNames[i]})
		}
	}
	// 6. Llamar a makeAllSpeak con tu slice
	for i := 0; i < numSpeakers; i++ {
		orator := orators[i]
		makeAllSpeak(orator)
	}

}
