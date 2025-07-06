package concorrencia

import "fmt"

func foo[T any](arg T) {
	fmt.Println(arg)
}

func bar() {
	foo("Hello, Generics!")
	foo(42)
	foo(3.14)
	foo([]int{1, 2, 3})
	foo(map[string]int{"one": 1, "two": 2})
	foo(struct {
		Name string
		Age  int
	}{"Alice", 30})
}
