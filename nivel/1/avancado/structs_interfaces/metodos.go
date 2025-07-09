package structsinterfaces

import "fmt"

func (u User) Foo() {
	fmt.Println(u.Name)
}

func (u *User) UpdateName(newName string) {
	u.Name = newName
}

func InterageMetodo() {
	user := User{Name: "John Doe"}
	user.Foo()
	user.UpdateName("Jane Doe")
	user.Foo()
}
