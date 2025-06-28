package structsinterfaces

import "fmt"

type Person struct {
	Name string
}

func (p Person) String() string {
	return "Person: " + p.Name
}

func TesteStringer() {
	p := Person{Name: "John"}
	fmt.Println(p.String())
}
