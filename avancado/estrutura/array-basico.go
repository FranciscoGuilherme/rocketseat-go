package estrutura

import "fmt"

func ArrayBasico() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	arr[2] = 15
	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)
}
