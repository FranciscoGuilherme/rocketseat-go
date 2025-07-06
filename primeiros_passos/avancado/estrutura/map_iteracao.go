package estrutura

import "fmt"

func MapIteracao() {
	newMap := map[string]string{
		"João":     "Joãozinho",
		"Maria":    "Mariquinha",
		"Pedro":    "Pedrinho",
		"Lucas":    "Luquinha",
		"Fernanda": "Nanda",
	}

	for key, value := range newMap {
		fmt.Printf("Chave: %s, Valor: %s\n", key, value)
	}
}
