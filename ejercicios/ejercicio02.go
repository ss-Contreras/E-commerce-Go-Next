package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func TablasNumeros() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese un numero: ")
	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())

		} else {
			for i := 0; i < 10; i++ {
				fmt.Println(numero, "* ", i, " = ", numero*i)
			}

		}
	}
}
