package main

import "fmt"

// updateQuantity actualiza la cantidad de un artículo y devuelve la nueva cantidad
// y si el artículo está agotado (cantidad = 0)
func updateQuantity(quantity *int, change int) (int, bool) {
    // TODO: Actualizar el puntero de cantidad y devolver el nuevo valor y el estado de agotado
	state := false
	*quantity = *quantity + change
	if *quantity < 0 {
		state = true
		*quantity = 0
	}
    return *quantity, state
}

// calculateValue calcula el valor total de un artículo basándose en la cantidad y el precio
// También actualiza mostValuableItem si este artículo es más valioso
func calculateValue(itemName string, quantity int, price float64, mostValuableItem *string, highestValue *float64) float64 {
    // TODO: Calcular el valor total, actualizar mostValuableItem si es necesario y devolver el valor total
	// switch itemName {
	// case "Apples":
	// 	price := 0.5
	// case "Oranges":
	// 	price := 0.7
	// case "Bananas":
	// 	price := 0.3
	// }
	totalValue := float64(quantity) * price
	if totalValue > *highestValue {
		*mostValuableItem = itemName
		*highestValue = totalValue
	}
    return totalValue
}

// displayInventory imprime los detalles del inventario
// Toma punteros para permitir que la función muestre datos en tiempo real
func displayInventory(apples, oranges, bananas *int, applePrice, orangePrice, bananaPrice float64) {
    fmt.Printf("Apples: %d (Value: $%.2f)\n", *apples, float64(*apples) * applePrice)
    fmt.Printf("Oranges: %d (Value: $%.2f)\n", *oranges, float64(*oranges) * orangePrice)
    fmt.Printf("Bananas: %d (Value: $%.2f)\n", *bananas, float64(*bananas) * bananaPrice)
}

func main() {
    // Inicializar inventario
    apples := 10
    oranges := 15
    bananas := 8
    
    // Precios
    applePrice := 0.5  // $0.50 cada uno
    orangePrice := 0.7 // $0.70 cada uno
    bananaPrice := 0.3 // $0.30 cada uno
    
    // Rastrear el artículo más valioso
    var mostValuableItem string
    var highestValue float64
    
    // Mostrar inventario inicial
    fmt.Println("Initial Inventory:")
    displayInventory(&apples, &oranges, &bananas, applePrice, orangePrice, bananaPrice)
    
    // Calcular valores iniciales y encontrar el artículo más valioso
    appleValue := calculateValue("Apples", apples, applePrice, &mostValuableItem, &highestValue)
    orangeValue := calculateValue("Oranges", oranges, orangePrice, &mostValuableItem, &highestValue)
    bananaValue := calculateValue("Bananas", bananas, bananaPrice, &mostValuableItem, &highestValue)
    
    fmt.Printf("Total inventory value: $%.2f\n", appleValue+orangeValue+bananaValue)
    fmt.Printf("Most valuable item: %s\n\n", mostValuableItem)
    
    // Simular algunas ventas
    fmt.Println("Processing sales...")
    _, applesOutOfStock := updateQuantity(&apples, -4) // Vender 4 manzanas
    _, orangesOutOfStock := updateQuantity(&oranges, -8) // Vender 8 naranjas
    _, bananasOutOfStock := updateQuantity(&bananas, -10) // Intentar vender 10 plátanos (más de los que tenemos)
    
    // Comprobar si algún artículo está agotado
    if applesOutOfStock {
        fmt.Println("Apples are out of stock!")
    }
    if orangesOutOfStock {
        fmt.Println("Oranges are out of stock!")
    }
    if bananasOutOfStock {
        fmt.Println("Bananas are out of stock!")
    }
    
    // Mostrar inventario actualizado
    fmt.Println("\nUpdated Inventory:")
    displayInventory(&apples, &oranges, &bananas, applePrice, orangePrice, bananaPrice)
    
    // Restablecer el seguimiento del más valioso para la recalculación
    mostValuableItem = ""
    highestValue = 0
    
    // Recalcular valores
    appleValue = calculateValue("Apples", apples, applePrice, &mostValuableItem, &highestValue)
    orangeValue = calculateValue("Oranges", oranges, orangePrice, &mostValuableItem, &highestValue)
    bananaValue = calculateValue("Bananas", bananas, bananaPrice, &mostValuableItem, &highestValue)
    
    fmt.Printf("Total inventory value: $%.2f\n", appleValue+orangeValue+bananaValue)
    fmt.Printf("Most valuable item: %s\n\n", mostValuableItem)
    
    // Reabastecer artículos
    fmt.Println("Restocking...")
    updateQuantity(&apples, 5)  // Añadir 5 manzanas
    updateQuantity(&oranges, 10) // Añadir 10 naranjas
    updateQuantity(&bananas, 12) // Añadir 12 plátanos
    
    // Mostrar inventario final
    fmt.Println("\nFinal Inventory:")
    displayInventory(&apples, &oranges, &bananas, applePrice, orangePrice, bananaPrice)
    
    // Restablecer el seguimiento del más valioso para el cálculo final
    mostValuableItem = ""
    highestValue = 0
    
    // Cálculo de valor final
    appleValue = calculateValue("Apples", apples, applePrice, &mostValuableItem, &highestValue)
    orangeValue = calculateValue("Oranges", oranges, orangePrice, &mostValuableItem, &highestValue)
    bananaValue = calculateValue("Bananas", bananas, bananaPrice, &mostValuableItem, &highestValue)
    
    fmt.Printf("Total inventory value: $%.2f\n", appleValue+orangeValue+bananaValue)
    fmt.Printf("Most valuable item: %s\n", mostValuableItem)
}