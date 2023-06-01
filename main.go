package main

import (
	"fmt"

	"github.com/ss-Contreras/E-commerce-Go-Next/variables"
)

func main() {
	variables.MuestroEnteros()
	fmt.Println("numeros Resto")
	variables.RestoVariables()
	fmt.Println("===========")
	estado, texto := variables.ConviertoaTexto(1222)
	fmt.Println(estado)
	fmt.Println(texto)

}
