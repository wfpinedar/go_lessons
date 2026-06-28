package main

import "fmt"

// addItems añade nuevos artículos y sus precios a la lista de compras
func addItems(names []string, prices []float64, newNames []string, newPrices []float64) ([]string, []float64) {
	for index, value := range newNames {
		names = append(names, value)
		prices = append(prices, newPrices[index])
	}
	return names, prices
}

// removeItem elimina un artículo en el índice especificado
func removeItem(names []string, prices []float64, index int) ([]string, []float64) {
	names = append(names[:index], names[index+1:]...)
	prices = append(prices[:index], prices[index+1:]...)
	return names, prices
}

// findExpensiveItems devuelve artículos con precios por encima del umbral
func findExpensiveItems(names []string, prices []float64, threshold float64) []string {
	foundedItems := make([]string, len(prices))
	for index, value := range prices {
		if value > threshold {
			foundedItems = append(foundedItems, names[index])
		}
	}
	return foundedItems
}

// calculateTotalCost devuelve la suma de todos los precios
func calculateTotalCost(prices []float64) float64 {
	totalCost := 0.0
	for _, value := range prices {
    totalCost = totalCost + value
	}
	return totalCost
}

func main() {
	// Inicializar listas de compras vacías
	names := []string{}
	prices := []float64{}
	
	// Añadir artículos iniciales
	initialNames := []string{"Apples", "Milk", "Bread"}
	initialPrices := []float64{2.99, 3.49, 2.29}
	
	names, prices = addItems(names, prices, initialNames, initialPrices)
	
	// Imprimir la lista de compras inicial
	fmt.Println("Initial Shopping List:")
	for i := range names {
		fmt.Printf("%d. %s - $%.2f\n", i, names[i], prices[i])
	}
	
	// Calcular e imprimir el costo total
	total := calculateTotalCost(prices)
	fmt.Printf("\nTotal Cost: $%.2f\n", total)
	
	// Encontrar e imprimir artículos caros
	priceThreshold := 3.00
	expensiveItems := findExpensiveItems(names, prices, priceThreshold)
	
	fmt.Printf("\nExpensive Items (above $%.2f):\n", priceThreshold)
	for _, item := range expensiveItems {
		fmt.Println(item)
	}
	
	// Añadir un nuevo artículo
	names, prices = addItems(names, prices, []string{"Cheese"}, []float64{4.99})
	
	// Eliminar un artículo
	names, prices = removeItem(names, prices, 1)
	fmt.Println("\nRemoved item at index 1")
	
	// Imprimir la lista de compras final
	fmt.Println("\nFinal Shopping List:")
	for i := range names {
		fmt.Printf("%d. %s - $%.2f\n", i, names[i], prices[i])
	}
}
