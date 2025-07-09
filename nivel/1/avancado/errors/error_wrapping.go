package erros

import (
	"errors"
	"fmt"
)

func foo() error {
	err := bar()
	if err != nil {
		//return errors.New("due erro em foo" + err.Error())
		return fmt.Errorf("erro em foo: %w", err)
	}
	return nil
}

var ErrQualquer = errors.New("erro qualquer")

func bar() error {
	return ErrQualquer
}

func ErrosWrappingExample() {
	err := foo()
	if err != nil && errors.Is(err, ErrQualquer) {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Tudo certo!")
}
