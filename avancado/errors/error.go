package erros

import (
	"errors"
	"fmt"
)

var ErrDivisaoPorZero error = errors.New("divisão por zero")

func ErrosExample() {
	resultado, err := dividir(10, 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Resultado:", resultado)
}

func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisaoPorZero
	}
	return a / b, nil
}
