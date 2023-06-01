package main

import (
	"fmt"
	"runtime"

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
	if os := runtime.GOOS; os == "Linux." || os == "OS X" {
		fmt.Println("Esto no es windows")
	} else {
		fmt.Println("Esto ess windows")
	}

}
