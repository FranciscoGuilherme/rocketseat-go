package modulo1

import (
	"fmt"
	"rocketseat-go/nivel/1/fundamentos/loops"
	"rocketseat-go/nivel/1/fundamentos/variaveis"
	"rocketseat-go/nivel/1/fundamentos/funcoes"
	"rocketseat-go/nivel/1/fundamentos/deferpackage"
	"rocketseat-go/nivel/1/fundamentos/variaveis/tipos"
)

func Modulo1Main() {
	sobreFuncoes()
	sobreVariaveis()
	sobreTipos()
	loops.LoopFor()
	loops.LoopParalelismo()
	deferpackage.DoDefer()
	deferpackage.DoDeferVariable()
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
	var nome, sobrenome, cidade, endereco string = variaveis.Variaveis()

	fmt.Println("Nome:", nome, "Sobrenome:", sobrenome)
	fmt.Println("Cidade:", cidade, "Endereço:", endereco)
}

func sobreTipos() {
	numero, converteNumero := tipos.TipoConversao()
	arr, another := tipos.ManipulaArray()

	fmt.Println("Bool:", tipos.TipoBool())
	fmt.Println("Número:", numero)
	fmt.Println("Número convertido para float64:", converteNumero)
	fmt.Println("Array:", arr)
	fmt.Println("Outro Array:", another)
}
