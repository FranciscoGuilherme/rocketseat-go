package erros

import "errors"

var (
	err1 = errors.New("erro 1")
	err2 = errors.New("erro 2")
)

func foo2() error {
	return err1
}

func bar2() error {
	return err2
}

func ErrorJoinExample() error {
	var errorsList error
	if err := foo2(); err != nil {
		errorsList = errors.Join(errorsList, err)
	}
	if err := bar2(); err != nil {
		errorsList = errors.Join(errorsList, err)
	}
	return errorsList
}
