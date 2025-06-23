package variaveis

import "fmt"

var idade int = 30

func Exemplo() {
	var nome, sobrenome string = "João", "Silva"
	var (
		cidade   string = "São Paulo"
		endereco string = "Rua A, 123"
	)
	altura := 1.75

	fmt.Println("Nome:", nome, "Sobrenome:", sobrenome, "Idade:", idade)
	fmt.Println("Cidade:", cidade, "Endereço:", endereco)
	fmt.Println("Altura:", altura)
}
