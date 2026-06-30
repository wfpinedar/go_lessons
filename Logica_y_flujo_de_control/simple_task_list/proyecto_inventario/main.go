package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Product struct {
	Price    float64
	Quantity int
}

func main() {
	var inventoryData, operations, parameters string
	fmt.Scanln(&inventoryData)
	fmt.Scanln(&operations)
	fmt.Scanln(&parameters)

	// Inicializar inventario
	inventory := make(map[string]Product)
	if inventoryData != "" {
		items := strings.Split(inventoryData, ",")
		for _, item := range items {
			parts := strings.Split(item, ":")
			if len(parts) == 3 {
				name := parts[0]
				price, _ := strconv.ParseFloat(parts[1], 64)
				quantity, _ := strconv.Atoi(parts[2])
				inventory[name] = Product{Price: price, Quantity: quantity}
			}
		}
	}

	// Mostrar mensaje de inicio
	fmt.Println("=== INVENTORY MANAGEMENT SYSTEM ===")
	fmt.Printf("System initialized with %d products\n", len(inventory))
	fmt.Println("Starting interactive session...")

	// Analizar operaciones y parámetros
	ops := strings.Split(operations, ",")
	params := strings.Split(parameters, "|")

	// Ejecutar cada operación
	for i, op := range ops {
		var param string
		if i < len(params) {
			param = params[i]
		}

		switch op {
		case "check":
			fmt.Println("--- STOCK CHECK ---")
			checkOperation(inventory, param)
			fmt.Println("Operation completed. Continuing to next operation...")

		case "add":
			fmt.Println("--- ADD ITEM ---")
			addOperation(inventory, param)
			fmt.Println("Operation completed. Continuing to next operation...")

		case "update":
			fmt.Println("--- UPDATE STOCK ---")
			updateOperation(inventory, param)
			fmt.Println("Operation completed. Continuing to next operation...")

		case "report":
			fmt.Println("--- GENERATE REPORT ---")
			reportOperation(inventory, param)
			fmt.Println("Operation completed. Continuing to next operation...")

		case "exit":
			fmt.Println("--- SYSTEM EXIT ---")
			exitOperation(inventory)
		}
	}
}

func checkStock(inventory map[string]Product, productName string) (int, error) {
	product, exists := inventory[productName]
	if !exists {
		return 0, errors.New("product not found")
	}
	return product.Quantity, nil
}

func checkOperation(inventory map[string]Product, param string) {
	productName := param
	fmt.Printf("Checking stock for: %s\n", productName)

	quantity, err := checkStock(inventory, productName)
	if err != nil {
		fmt.Println("Product not found in inventory")
	} else {
		fmt.Printf("Stock level: %d units\n", quantity)
	}
}

func addNewItem(inventory map[string]Product, name string, price float64, quantity int) error {
	if _, exists := inventory[name]; exists {
		return errors.New("product already exists")
	}
	inventory[name] = Product{Price: price, Quantity: quantity}
	return nil
}

func addOperation(inventory map[string]Product, param string) {
	parts := strings.Split(param, ":")
	if len(parts) != 3 {
		fmt.Println("Error adding product: invalid format")
		return
	}

	name := parts[0]
	price, _ := strconv.ParseFloat(parts[1], 64)
	quantity, _ := strconv.Atoi(parts[2])

	fmt.Printf("Adding new product: %s\n", name)

	err := addNewItem(inventory, name, price, quantity)
	if err != nil {
		fmt.Printf("Error adding product: %s\n", err.Error())
	} else {
		fmt.Println("Product added successfully")
	}
}

func updateStock(inventory map[string]Product, productName string, change int) error {
	product, exists := inventory[productName]
	if !exists {
		return errors.New("product not found")
	}

	newQuantity := product.Quantity + change
	if newQuantity < 0 {
		return errors.New("insufficient stock")
	}

	product.Quantity = newQuantity
	inventory[productName] = product
	return nil
}

func updateOperation(inventory map[string]Product, param string) {
	parts := strings.Split(param, ":")
	if len(parts) != 2 {
		fmt.Println("Update error: invalid format")
		return
	}

	productName := parts[0]
	change, _ := strconv.Atoi(parts[1])

	fmt.Printf("Updating stock for: %s\n", productName)

	err := updateStock(inventory, productName, change)
	if err != nil {
		fmt.Printf("Update error: %s\n", err.Error())
	} else {
		newQuantity := inventory[productName].Quantity
		if change >= 0 {
			fmt.Printf("Added %d units. New stock: %d\n", change, newQuantity)
		} else {
			fmt.Printf("Removed %d units. New stock: %d\n", -change, newQuantity)
		}
	}
}

func generateReport(inventory map[string]Product, reportType string, threshold int) {
	// Obtener nombres ordenados
	var names []string
	for name := range inventory {
		names = append(names, name)
	}
	sort.Strings(names)

	if reportType == "full" {
		fmt.Println("=== FULL REPORT ===")
		fmt.Printf("Complete inventory (minimum threshold: %d):\n", threshold)

		for _, name := range names {
			product := inventory[name]
			status := "[OK]"
			if product.Quantity < threshold {
				status = "[LOW STOCK]"
			}
			fmt.Printf("- %s: %d units @ $%.2f each %s\n",
				name, product.Quantity, product.Price, status)
		}
	} else if reportType == "low" {
		fmt.Println("=== LOW REPORT ===")
		fmt.Printf("Products with stock below %d:\n", threshold)

		hasLowStock := false
		for _, name := range names {
			product := inventory[name]
			if product.Quantity < threshold {
				hasLowStock = true
				fmt.Printf("- %s: %d units (Price: $%.2f)\n",
					name, product.Quantity, product.Price)
			}
		}

		if !hasLowStock {
			fmt.Println("No products below threshold")
		}
	}
}

func reportOperation(inventory map[string]Product, param string) {
	parts := strings.Split(param, ",")
	if len(parts) != 2 {
		fmt.Println("Error generating report: invalid format")
		return
	}

	reportType := parts[0]
	threshold, _ := strconv.Atoi(parts[1])

	fmt.Printf("Generating %s report with threshold %d\n", reportType, threshold)
	generateReport(inventory, reportType, threshold)
}

func exitOperation(inventory map[string]Product) {
	totalProducts := len(inventory)
	totalItems := 0
	totalValue := 0.0

	for _, product := range inventory {
		totalItems += product.Quantity
		totalValue += float64(product.Quantity) * product.Price
	}

	fmt.Println("Final inventory status:")
	fmt.Printf("Total products: %d\n", totalProducts)
	fmt.Printf("Total items: %d\n", totalItems)
	fmt.Printf("Total value: $%.2f\n", totalValue)
	fmt.Println("Session completed successfully")
	fmt.Println("Thank you for using the Inventory Management System")
}
