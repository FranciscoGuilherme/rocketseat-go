package estrutura

import "fmt"

func Map() {
	var m map[string]string = make(map[string]string, 100)
	newMap := map[string]string{
		"João": "Joãozinho",
		"Maria": "Mariquinha",
	}
	fmt.Println(m)
	fmt.Println(newMap)
}
