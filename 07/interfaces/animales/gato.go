package animales

import "fmt"

type Gato struct {
	Nombre string
}

func (g Gato) Comunicarse() {
	fmt.Println("El gato esta maullando")
}
