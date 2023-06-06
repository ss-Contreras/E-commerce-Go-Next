package ejercicios

import (
	"strconv"
)

func RetornarValor(texto string) (int, error) {
	num, err := strconv.Atoi(texto)
	if err != nil {
		return 0, err
	}
	if num > 100 {
		return num, nil
	} else {
		return num, nil
	}
}
