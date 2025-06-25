package deferpackage

import "fmt"

func DoDefer() {
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 1")
	defer fmt.Println("world")
	fmt.Println("hello")
}

func DoDeferVariable() {
	var x int = 10
	defer func() {
		fmt.Println("defer variable x =", x)
	}()

	x = 20
	fmt.Println("x =", x)
}
