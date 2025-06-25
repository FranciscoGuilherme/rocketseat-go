package loops

import "fmt"

func LoopFor() {
	var arr [3]int = [3]int{1, 2, 3}

	for i := 0; i < 10; i++ {}
	for _, value := range arr {
		fmt.Println(value)
	}
	for index := range 10 {
		fmt.Println("Looping sem variável", index)
	}
}
