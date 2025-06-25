package condicionais

import "fmt"

func CondicionalSwitch(x int) {
	switch x {
		case 1:
			fmt.Println("O valor é 1")
			fallthrough
		case 2: fmt.Println("O valor é 2")
		case 3: fmt.Println("O valor é 3")
		default: fmt.Println("O valor não é 1, 2 ou 3")
	}

	switch {
		case "abc" == "foo": fmt.Println("abc é igual a foo")
	}
}
