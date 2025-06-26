package funcoes

func FuncaoClosure(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}
