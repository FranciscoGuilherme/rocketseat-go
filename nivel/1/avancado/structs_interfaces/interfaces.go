package structsinterfaces

import "fmt"

type Animal interface {
	Sound() string
}

type Dog struct{}

func (d Dog) Sound() string {
	return "Au! Au!"
}

func WhatDoesTheAnimalSay(a Animal) string {
	return a.Sound()
}

func InterfaceExample() {
	var dog Dog = Dog{}
	fmt.Println("O cachorro diz:", WhatDoesTheAnimalSay(dog))
}
