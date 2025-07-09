package estrutura

import "fmt"

func MapSlices() {
	m := map[string][]int{
		"Pedro": {1, 2, 3},
	}

	fmt.Println(m)
}
