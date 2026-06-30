package main

import "fmt"

type Player struct {
	Name string
}

// TODO: Define tu interfaz MediaPlayer aquí
type MediaPlayer interface {
	Play() string
	Pause() string
	Stop() string
}

type SimplePlayer struct{}

// Play implementa el método Play de la interfaz MediaPlayer.
func (sp SimplePlayer) Play() string {
	return "Play"
}

// Pause implementa el método Pause de la interfaz MediaPlayer.
func (sp SimplePlayer) Pause() string {
	return "Pause"
}

// Stop implementa el método Stop de la interfaz MediaPlayer.
func (sp SimplePlayer) Stop() string {
	return "Stop"
}

func main() {
	// Leer entrada
	var action string
	fmt.Scanln(&action)

	player := SimplePlayer{}

	// TODO: Escribe tu código a continuación para manejar la acción e imprimir la salida requerida
	switch action {
	case "play":
		fmt.Printf("MediaPlayer interface requires: %s() string\n", player.Play())
	case "pause":
		fmt.Printf("MediaPlayer interface requires: %s() string\n", player.Pause())
	case "stop":
		fmt.Printf("MediaPlayer interface requires: %s() string\n", player.Stop())
	default:
		fmt.Printf("Acción no válida: %s\n", action)
	}
	// Imprimir el resultado
	// Recuerda imprimir en formato: "MediaPlayer interface requires: [MethodName]() string"
}
