package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Read input
	var numProductsStr string
	var productDataStr string
	var operationsStr string

	fmt.Scanln(&numProductsStr)
	fmt.Scanln(&productDataStr)
	fmt.Scanln(&operationsStr)

	type Product struct {
		Price    float64
		Quantity int
	}

	// TODO: Create map to store product pointers
	products := make(map[string]*Product)
	product_list := strings.Split(productDataStr, ",")
	// operation_list := strings.Split(operationsStr, ",")
	for _, prod := range product_list {
		values := strings.Split(prod, ":")
		price, _ := strconv.ParseFloat(values[1], 64)
		quantity, _ := strconv.Atoi(values[2])
		products[values[0]] = &Product{price, quantity}
	}

	// TODO: Create slice to maintain product order (to ensure consistent output)

	// TODO: Parse product data and populate the map
	// Remember to:
	// - Split productDataStr by commas to get individual entries
	// - For each entry, split by colons to get name, price, quantity
	// - Convert price and quantity strings to appropriate types
	// - Store pointer to Product struct in map
	// - Add product name to order slice

	// TODO: Display initial inventory
	// Use the order slice to iterate through products consistently
	// Format: "[name]: $[price] (Stock: [quantity])"

	// TODO: Parse and apply update operations
	// Remember to:
	// - Split operationsStr by commas to get individual operations
	// - For each operation, split by colons to get type, name, value
	// - Update the appropriate field directly through the map pointer
	// - Print update message for each operation

	// TODO: Display updated inventory
	// Again, use the order slice for consistent output
	// Calculate total value while displaying

	// TODO: Calculate and display total inventory value
	// Format: "Total Inventory Value: $[total_value]"
}
