package main

import (
	"fmt"
)
func main (){
	//arrays
	//todo los valores del mismo tipo
	//tamaño fijo
	/*var names [10]string
	names[1] = "Bryan"
	names[2]= "augusto"*/

	names :=[3] string{"bryan","augusto","pedro"}
	//name := names[1]
	// array[inicio:final(excluyente)]
	fmt.Println(names[1:2])
}