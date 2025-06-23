package variaveis

var idade int = 30

func Variaveis() (string, string, string, string) {
	var nome, sobrenome string = "João", "Silva"
	var (
		cidade   string = "São Paulo"
		endereco string = "Rua A, 123"
	)

	return nome, sobrenome, cidade, endereco
}
