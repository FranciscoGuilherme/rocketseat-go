package main

import (
	"fmt"
	"MyModule/funcoes"
	"MyModule/variaveis"
	"MyModule/variaveis/tipos"
)

func main() {
	//sobreFuncoes()
	//sobreVariaveis()
	sobreTipos()
}

func sobreFuncoes() {
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

func sobreVariaveis() {
	variaveis.Exemplo()
}

func sobreTipos() {
	numero, converteNumero := tipos.TipoConversao()

	fmt.Println("Bool:", tipos.TipoBool())
	fmt.Println("Número:", numero)
	fmt.Println("Número convertido para float64:", converteNumero)
}
