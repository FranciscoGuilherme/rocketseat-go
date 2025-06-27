package structsinterfaces

import "fmt"

type MinhaString string
type User struct {
	ID    uint64
	Age   int
	Name  string
	Email string
}

func Thing() {
	user := User{
		ID:    1,
		Age:   30,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	fmt.Println(user)
}
