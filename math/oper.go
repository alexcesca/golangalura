package math

import "errors"

// Soma é uma função de retorna a soma de dois numeros inteiros.
func Soma(a int, b int) int {
	return a + b
}

func Multiplica(a int, b int) (int, error) {
	res := a * b
	if (a == 0) || (b == 0) {
		return 0, errors.New("Retorno zerado")
	}

	return res, nil
}
