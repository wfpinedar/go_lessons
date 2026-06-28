package main

import (
	"fmt"
	"strconv"
)

func processNotification(data interface{}) {
	switch v := data.(type) {
	case int:
		fmt.Printf("SMS notification with %d characters", v)
	case string:
		fmt.Printf("Email notification: %s\n", v)
	case bool:
		fmt.Printf("Push notifications: %t", v)
	case float64:
		fmt.Printf("Alert with priority: %g", v)
	default:
		fmt.Println("Unknown notification type")
	}
}

func main() {
	// Leer entrada
	var notificationType string
	var content string
	fmt.Scanln(&notificationType)
	fmt.Scanln(&content)

	// TODO: Escribe tu código a continuación
	// 1. Crea la función processNotification que toma un parámetro interface{}

	var value interface{}
	// 2. Convierte el contenido al tipo apropiado según notificationType
	switch notificationType {
	case "int":
		if num, err := strconv.Atoi(content); err == nil {
			value = num
		}
	case "float64":
		if num, err := strconv.ParseFloat(content, 64); err == nil {
			value = num
		}
	case "string":
		value = content
	case "bool":
		value = content == "true"
	}
	// 3. Llama a processNotification con el valor convertido
	processNotification(value)
	// 4. Usa un type switch dentro de processNotification para manejar diferentes tipos

}
