package main

import (
	"fmt"
	"strings"
)

// basic function
func hi() {
	fmt.Println("Hi")
}

// function with parameters
func birthday(month string, day int) {
	fmt.Printf("My birthday is on the %dth of %s\n", day, month)
}

// return names
func division(a, b int) (cosciente, resto int) {
	cosciente = a / b
	resto = a % b
	return // <- implicit return here
}


// variadic function
func greet(yo string, personas ...string) {
	fmt.Printf("Hola %s, soy %s. \n", strings.Join(personas, ", "), yo)
}

// recursive function
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

// closures functions
func saludar(yo string) func(persona string) {
	return func(persona string) {
		fmt.Printf("Hola %s, me llamo %s\n", persona, yo)
	}
}

// defer function
func example_defer() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}

type Persona struct {
	nombre string
	edad   int
}

// receiver function
func (p *Persona) GetNombre() string {
	return p.nombre
}

func main() {
	greet("thiago", "jose", "pedro", "juan", "antonio")

	fmt.Println(factorial(5))
	fmt.Printf("%d! = %d \n", 4, factorial(4))
	fmt.Printf("%d! = %d \n", 5, factorial(5))
	fmt.Printf("%d! = %d \n", 7, factorial(7))
	fmt.Printf("%d! = %d \n", 0, factorial(0))

	// IIFE expression
	func(name string) {
		fmt.Printf("Hi %s!\n", name)
	}("Thiago")

	// anonymous function
	sum := func(a, b int) int {
		return a + b
	}

	fmt.Println(sum(1, 2))

	ThiagoSaludo := saludar("Thiago")
	ThiagoSaludo("Jose")

	birthday("April", 25)

	hi()

	division(10, 5)

	example_defer()

}

division()
