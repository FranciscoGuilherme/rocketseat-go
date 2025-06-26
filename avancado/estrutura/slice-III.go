package estrutura

import "fmt"

func SliceIII() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	slice := arr[:5:5]
	fmt.Println(slice)
}

func SlicePerformance(slice []int) {
	_ = slice[3]
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	fmt.Println(slice[3])
}
