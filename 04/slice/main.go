package main

import "fmt"

func main() {

	//slices
	// 1. apuntador a un array
	// 2. tamaño
	// 3. capacidad

	//var name []string
	// M A K E(tipode dato del slice, tamaño inicial, capacidad inicial)
	// 
	//name := make( []string,0 ) // la forma mas comun de realizar la creacion de un slice

	name:= []string{
		"Bryan",
		"Augusto",
		"Antonio",
	}


	/*
	fmt.Printf("Su tamaño es : %d y su capacidad es: %d\n", len(name), cap(name))
	name = append(name, "bryan")
	fmt.Printf("Su tamaño es : %d y su capacidad es: %d\n", len(name), cap(name))
	name = append(name, "augusto")
	fmt.Printf("Su tamaño es : %d y su capacidad es: %d\n", len(name), cap(name))
	name = append(name, "Carmin")
	fmt.Printf("Su tamaño es : %d y su capacidad es: %d\n", len(name), cap(name))
	name = append(name, "Fernando")
	fmt.Printf("Su tamaño es : %d y su capacidad es: %d\n", len(name), cap(name))
	name = append(name, "Diego")
	fmt.Printf("Su tamaño es : %d y su capacidad es: %d\n", len(name), cap(name))
	name = append(name, "Maria")
	fmt.Printf("Su tamaño es : %d y su capacidad es: %d\n", len(name), cap(name))
	name = append(name, "Eduardo")
	fmt.Printf("Su tamaño es : %d y su capacidad es: %d\n", len(name), cap(name))

	name[3] = "Antonio"
*/
	fmt.Println(name)
}
