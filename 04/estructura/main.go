package main

import "fmt"

// Persona es una estructura
type Persona struct {
	edad      uint8
	name      string
	firstName string
	email []string
	comida map[string]uint8
}

func main() {
	// estructuras es equivalente a clases en otros lenguajes
	// las estructuras son tipos de datos
	/*
	
	var p1 Persona
	p1.name = "Bryan Augusto"
	p1.edad = 30
	p1.firstName = "Cruz Castro"
	
	*/
	comida := map[string]uint8 {"mango":10, "manzana":5}
	p1 := Persona{
		12,
		"Augusto",
		 "Castro",
		 []string{"a@gmail.com", "b@outlook.es"},
		 comida,
		}
	

	fmt.Println(p1)
}
