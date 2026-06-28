package main

import (
	"fmt"
	"sort"
)

// addWord añade una palabra al mapa contador
func addWord(counter map[string]int, word string) {
	if _, exists := counter[word]; exists {
		counter[word]++
	} else {
		counter[word] = 1
	}
}

// getCount devuelve el recuento de una palabra específica
func getCount(counter map[string]int, word string) int {
	count := counter[word] 
	return count

}

// removeWord elimina una palabra del contador
// Devuelve verdadero si la palabra fue encontrada y eliminada
func removeWord(counter map[string]int, word string) bool {
	_, exists := counter[word]
	if exists {delete(counter, word)}
	return exists
}

// printWordCounts imprime los recuentos de palabras en orden alfabético
func printWordCounts(counter map[string]int) {
	// Obtener todas las claves (palabras) del mapa
	words := make([]string, 0, len(counter))
	for word := range counter {
		words = append(words, word)
	}
	
	// Ordenar las palabras alfabéticamente
	sort.Strings(words)
	
	// Imprimir cada palabra y su recuento en orden ordenado
	for _, word := range words {
		fmt.Printf("%s: %d\n", word, counter[word])
	}
}

func main() {
	// Inicializar el mapa contador de palabras
	wordCounter := make(map[string]int)
	
	// Añadir algunas palabras
	words := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
	for _, word := range words {
		addWord(wordCounter, word)
	}
	
	// Imprimir los recuentos de palabras
	fmt.Println("Recuentos de palabras:")
	printWordCounts(wordCounter)
	
	// Obtener el recuento de palabras específicas
	fmt.Printf("\nRecuento de 'apple': %d\n", getCount(wordCounter, "apple"))
	fmt.Printf("Recuento de 'grape': %d\n", getCount(wordCounter, "grape"))
	
	// Eliminar una palabra
	removed := removeWord(wordCounter, "banana")
	fmt.Printf("\nEliminado 'banana': %v\n", removed)
	
	// Intentar eliminar una palabra que no existe
	removed = removeWord(wordCounter, "grape")
	fmt.Printf("Eliminado 'grape': %v\n", removed)
	
	// Imprimir los recuentos finales de palabras
	fmt.Println("\nRecuentos finales de palabras:")
	printWordCounts(wordCounter)
}