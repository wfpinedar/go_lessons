package main

import (
	"fmt"
)

// Counter es una estructura de datos de contador simple
type Counter struct {
	value int
	name  string
}

// NewCounter crea un nuevo contador con un nombre dado
func NewCounter(name string) Counter {
	return Counter{
		value: 0,
		name:  name,
	}
}

// Increment aumenta el contador en 1
func (c *Counter) Increment() {
	c.value++
	// Añadir código para incrementar el valor del contador
}

// Decrement disminuye el contador en 1 (pero no por debajo de 0)
func (c *Counter) Decrement() {
	// Añadir código para decrementar el valor del contador
	// Asegurarse de que el valor no baje de 0
	if c.value > 0 {
		c.value--
	}
}

// Reset restablece el contador a 0
func (c *Counter) Reset() {
	// Añadir código para restablecer el contador a 0
	c.value = 0
}

// Value devuelve el recuento actual
func (c Counter) Value() int {
	// Añadir código para devolver el valor actual
	return c.value // Reemplazar esto
}

// String devuelve una representación en cadena del contador
func (c Counter) String() string {
	return fmt.Sprintf("%s: %d", c.name, c.value)
}

func main() {
	// Crear un nuevo contador llamado "Visitors"
	visitors := NewCounter("Visitors")
	
	// Incrementar el contador varias veces
	visitors.Increment()
	visitors.Increment()
	visitors.Increment()
	
	// Imprimir el recuento actual
	fmt.Println(visitors)
	fmt.Println("Current value:", visitors.Value())
	
	// Decrementar el contador
	visitors.Decrement()
	fmt.Println(visitors)
	
	// Restablecer el contador
	visitors.Reset()
	fmt.Println(visitors)
	
	// Probar que el contador no baja de cero
	visitors.Decrement()
	fmt.Println(visitors)
}