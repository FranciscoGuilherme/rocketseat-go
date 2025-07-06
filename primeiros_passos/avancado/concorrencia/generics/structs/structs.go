package main

import "fmt"

type NovoTipo[T any] struct {
	Vari T
}

func (n NovoTipo[T]) FuncCustom() {}

func main() {
	var novoTipo NovoTipo[string] = NovoTipo[string]{}
	fmt.Println(novoTipo)
}
