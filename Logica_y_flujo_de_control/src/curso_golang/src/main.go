package main

import "fmt"

func main() {

	switch modulo := 6 % 2; modulo {
	case 0:
		fmt.Println("Even")
	case 1:
		fmt.Println("Odd")
	default:
		fmt.Println("Unknown")
	}

	// Sin condicional switch
	value := 99
	switch {
	case value > 100:
		fmt.Println("mayor a 100")
	case value < 0:
		fmt.Println("menor a 0")
	default:
		fmt.Println("No condition")
	}
}
