package main

import (
	"fmt"
)

func main() {

	a() // -> 0

	b() // -> 3 2 1 0

	result := c()
	fmt.Println(result) // -> 2
}

// DEFER RULES

// 1. A deferred function’s arguments are evaluated when the defer statement is evaluated.
// 1. Los argumentos de una funcion diferida se evaluan cuando se evalua la declaracion defer.

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	// fmt.Println(i) -> nos daria como resultado 1, sin la llamada al defer.
	return
}

// 2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
// 2. Las llamadas a funciones diferidas se ejecutan en el orden: Último en entrar, primero en salir; después de que regresa la función return.

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

// 3. Deferred functions may read and assign to the returning function’s named return values.
// 3. Las funciones defer pueden leer y asignar a los valores de retorno nombrados de la función que regresa.

func c() (i int) {
	defer func() { i++ }()
	return 1
}
