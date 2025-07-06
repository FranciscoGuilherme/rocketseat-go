package memoria

import "fmt"

func ManipulandoPontiero() {
	x := 10
	p := &x

	fmt.Println("Valor de x:", x)
	fmt.Println("Valor de x:", *p)

	updateNumber(p)
	fmt.Println("Valor de x:", *p)
}

func updateNumber(n *int) {
	*n = 20
}

func create() *int {
	var x int = 10
	return &x
}
