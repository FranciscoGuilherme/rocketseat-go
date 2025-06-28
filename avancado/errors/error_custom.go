package erros

import "fmt"

type SqrtError struct {
	Message string
}

func (e SqrtError) Error() string {
	return e.Message
}

func raizQuadrada(x float64) (float64, error) {
	if x < 0 {
		return 0, SqrtError{
			Message: "não é possível calcular a raiz quadrada de um número negativo",
		}
	}
	return x * x, nil
}

func ErrosCustomizadosExample() {
	resultado, err := raizQuadrada(-10)
	if err != nil {
		fmt.Println("Erro:", err.Error())
		return
	}
	fmt.Println("Resultado:", resultado)
}
