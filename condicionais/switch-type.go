package condicionais

import "fmt"

func CondicionalType(x any) {
	switch x.(type) {
		case int:
			fmt.Println("O valor é um inteiro")
		case string:
			fmt.Println("O valor é uma string")
		case bool:
			fmt.Println("O valor é um booleano")
		default:
			fmt.Println("O valor é de um tipo desconhecido")
	}
}
