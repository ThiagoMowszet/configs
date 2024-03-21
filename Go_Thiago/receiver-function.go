package main

import "fmt"

type persona struct {
	nombre string
	edad   int8
	equipo string
}


/* var a = func() int {
    fmt.Println("a")
    return 0
}() */

func (p persona) valoresPersona() {
	fmt.Printf("%s, tiene %d y es hincha de %s hace mucho tiempo", p.nombre, p.edad, p.equipo)
}

func main() {
	Thiago := persona{
		nombre: "Thiago",
		edad:   22,
		equipo: "River",
	}

	// fmt.Print(Thiago)
	Thiago.valoresPersona()
}
