package jogo

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func Jogo() {
	fmt.Println("Jogo da adivinhação")
	fmt.Println("Escolha um número entre 1 e 100")

	var chutes [10]int64
	var numero int64 = rand.Int64N(101)
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

	for index := range chutes {
		fmt.Printf("Qual é o seu chute?")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("O seu chute não é um número válido.")
			return
		}

		switch {
			case chuteInt < numero:
				fmt.Println("O número é maior que o seu chute.", chuteInt)
			case chuteInt > numero:
				fmt.Println("O número é menor que o seu chute.", chuteInt)
			case chuteInt == numero:
				return
		}

		chutes[index] = chuteInt
	}

	fmt.Printf(
		"Você não acertou o número. O número era: %d\n" +
		"Essas foram as suas tentativas: %v\n", numero, chutes,
	)
}
