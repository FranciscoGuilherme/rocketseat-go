package structsinterfaces

import "fmt"

func AnyExample(a any) {
	str, ok := a.(string)
	fmt.Println(str, ok)
}
