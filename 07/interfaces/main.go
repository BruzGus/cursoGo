package main

import "github.com/BruzGus/cursoGo/07/interfaces/animales"


func main(){

	var p  animales.Perro
	var g  animales.Gato

	p.Nombre="Forrest"
	g.Nombre="Tigresa"

	// adoptarPerro(p)
	// adoptarGato(g)
	adoptarMascota(p)
	adoptarMascota(g)
}
/*
func adoptarPerro(p animales.Perro){
	p.Comunicarse()
}

func adoptarGato(g animales.Gato){
	g.Comunicarse()
	
}*/

func adoptarMascota(m animales.Mascota){
	m.Comunicarse()
}
