package estrutura

var filmesNoDB []string = []string{
	"Star Wars",
	"Matrix",
	"Vingadores",
	"Batman",
	"Superman",
	"Homem Aranha",
	"Capitão Marvel",
	"Thor",
	"Pantera Negra",
	"Guardiões da Galáxia",
}
var resultsFromApi []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func SliceII() []string {
	var filme string
	var filmes []string = make([]string, 0, len(resultsFromApi))

	for _, value := range resultsFromApi {
		filme = filmesNoDB[value]
		filmes = append(filmes, filme)
	}

	return filmes
}
