package main

import (
	"fmt"
	"strconv"
)

// TODO: Define la estructura Account aquí
type Account struct {
    Name string
    Balace float64
}
// TODO: Define el método deposit con receptor de puntero aquí
func (a *Account) deposit(depositAmount float64) {
    a.Balace = a.Balace + depositAmount
}
// TODO: Define el método withdraw con receptor de puntero aquí
func (a *Account) withdraw(withdrawAmount float64) {
    a.Balace = a.Balace - withdrawAmount
}
// TODO: Define el método displayBalance con receptor de valor aquí
func (a Account) displayBalance() {
    fmt.Printf("Account: %s, Balance: $%.2f\n", a.Name, a.Balace)
}

func main() {
	// Leer entrada
	var name string
	var initialBalanceStr string
	var transactionAmountStr string
	
	fmt.Scanln(&name)
	fmt.Scanln(&initialBalanceStr)
	fmt.Scanln(&transactionAmountStr)
	
	// Convertir entradas de cadena a tipos apropiados
	initialBalance, _ := strconv.ParseFloat(initialBalanceStr, 64)
	transactionAmount, _ := strconv.ParseFloat(transactionAmountStr, 64)
	
	// TODO: Crear una instancia de Account
	account := Account{name, initialBalance}
	// TODO: Mostrar saldo inicial, realizar depósito, mostrar saldo, realizar retiro, mostrar saldo final
  account.displayBalance()
  account.deposit(transactionAmount)
  account.displayBalance()
  account.withdraw(transactionAmount)
  account.displayBalance()
}
