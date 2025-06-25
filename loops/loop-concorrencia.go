package loops

import (
	"fmt"
	"sync"
)

func LoopParalelismo() {
	const tamanho = 10
	var wg sync.WaitGroup
	wg.Add(tamanho)

	for index := range tamanho {
		go func() {
			defer wg.Done()
			fmt.Println("Looping com variável", index)
		}()
	}

	wg.Wait()
}
