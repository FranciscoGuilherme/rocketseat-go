package main

type TipoEspecifico struct {}

type MinhaConstraint interface {
	Fazer()
}

func ComoFazer[T interface{ Fazer() }](arg T) {
	arg.Fazer()
}
