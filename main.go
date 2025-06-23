package main

import (
	"fmt"
	"MyModule/funcoes"
)

func main() {
	primeiro := 10
	segundo := 20
	divisao, resto := funcoes.FuncaoRetornos(10, 3)
	novoPrimeiro, novoSegundo := funcoes.FuncaoSwap(primeiro, segundo)

	fmt.Println("Hello world!")
	fmt.Println("Resto =", resto)
	fmt.Println("Divisão =", divisao)
	fmt.Println("Primeiro =", primeiro, "Segundo =", segundo)
	fmt.Println("Primeiro Swap =", novoPrimeiro, "Segundo Swap =", novoSegundo)
	fmt.Println("Soma com função", funcoes.FuncaoSomar(primeiro, segundo))
	fmt.Println("Soma com função closure =", funcoes.FuncaoClosure(10)(5))
	fmt.Println("Soma com função variática =", funcoes.FuncaoVariatica(1, 2, 3))
}
