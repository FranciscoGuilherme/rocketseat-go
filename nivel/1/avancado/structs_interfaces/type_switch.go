package structsinterfaces

import "fmt"

type Cat struct {}

func (c Cat) Sound() string {
	return "Meow!"
}

func SwitchExample(a Animal) {
	switch v := a.(type) {
	case Dog:
		fmt.Println("O cachorro diz:", v.Sound())
	case Cat:
		fmt.Println("O gato diz:", v.Sound())
	default:
		fmt.Println("Animal desconhecido")
	}
}
