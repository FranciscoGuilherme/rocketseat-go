package concorrencia

import "fmt"

type NovoTipo[T any] struct {
	Vari T
}

func (n NovoTipo[T]) FuncCustom() {}

func principal2() {
	var novoTipo NovoTipo[string] = NovoTipo[string]{}
	fmt.Println(novoTipo)
}
