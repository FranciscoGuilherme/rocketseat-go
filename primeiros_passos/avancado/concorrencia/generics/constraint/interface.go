package main

import "fmt"

type MeuTipo string

func (m MeuTipo) Foo() string {
	return "teste"
}

type MyConstraint interface {
	Foo() string
}

func fooConstraint[T MyConstraint](arg T) {
	fmt.Println("Argument:", arg.Foo())
}

func principal() {
	var mt MeuTipo = "Exemplo"
	fooConstraint(mt)
}
