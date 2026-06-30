package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

type Book struct {
    Title string
    Author string
    Pages int
}

func (b Book) displayInfo() {
    fmt.Printf("Title: %s, Author: %s, Pages: %d\n", b.Title, b.Author, b.Pages)
}

func (b Book) getDescription() string {
	message := b.Title + " by " + b.Author
	return message
}

func main() {
    // Crear escáner para leer líneas de entrada
    scanner := bufio.NewScanner(os.Stdin)
    
    // Leer título
    scanner.Scan()
    title := scanner.Text()
    
    // Leer autor
    scanner.Scan()
    author := scanner.Text()
    
    // Leer páginas
    scanner.Scan()
    pagesStr := scanner.Text()
    
    // Convertir la cadena de páginas a entero
    pages, _ := strconv.Atoi(pagesStr)
    
    // TODO: Crea una instancia de Book y llama a los métodos
    book := Book{title, author, pages}
	book.displayInfo()
	fmt.Println(book.getDescription())

}