package main

import (
	"strings"
    "strconv"
    "fmt"
)


func main() {

    strings.ToUpper("Thiago") // THIAGO
    strings.ToLower("THIAGO") // thiago
    strings.Count("alfajor", "a") // 2
    strings.Contains("pablito clavo un clavito cuantos clavitos clavo pablito", "pablito") // true
    strings.Split("Hola como estas?", ", ") // [Hola como estas?]
    // strings.Title("Thiago mowszet - software developer") // Deprecated: The rule Title uses for word boundaries does not handle Unicode
    strings.Trim("!!Hola gente¡¡", "!¡") // Hola gente
    strings.TrimSpace("     T")
    strings.EqualFold("Thiago", "THIago") //true -> EqualFold is not case sensitive
    strings.Index("Thiago Mowszet", "Mow") // 7
    strings.IndexAny("Thiago Mowszet", "o") // 5

    strings.Fields("Manzana, Naranaja, Frutilla") // [Manzana, Naranaja, Frutilla]
    strings.Join([]string {"Pera", "Manzana", "Banana"}, " - ")
    
    strings.Replace("hola, hola, hola amigos!", "hola", "adios", 1)  //adios, hola, hola amigos!"
	strings.Replace("hola, hola, hola amigos!", "hola", "adios", 2)  //"adios, adios, hola amigos!"
	strings.Replace("hola, hola, hola amigos!", "hola", "adios", -1) //"adios, adios, adios, amigos!"


    strings.Compare("A", "B") // -1
    strings.Compare("A", "A") // 0
    strings.Compare("B", "A") // 1

    strings.ContainsAny("ABCDEFGHI", "E") // false
    strings.ContainsAny("ABCDEFGHI", "GH") // true

    strings.HasPrefix("Hola a todos", "Ho") // true
    strings.HasSuffix("Hola a todos", "os") // true

    strings.Repeat("Thiago ", 10)
    conv()

}


func conv() {

    nombre := "Thiago"
    edad := "22"
    altura := "1.75"
    lesionado := "false"

    fmt.Printf("%T, %v\n", nombre, nombre) // string, Thiago

	if s, err := strconv.Atoi(edad); err == nil {
		fmt.Printf("%T, %v\n", s, s) // int, 22 
	}

	if s, err := strconv.ParseInt(edad, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s) //int64, 22 
	}

	if s, err := strconv.ParseFloat(altura, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s) // float64, 1.75
	}

	if s, err := strconv.ParseBool(lesionado); err == nil {
		fmt.Printf("%T, %v\n", s, s) // bool, false
	}
}

