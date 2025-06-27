package structsinterfaces

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Peso  int `json:"peso"`
	Marca string `json:"marca"`
}

func (c Car) PesoEmToneladas() float64 {
	return float64(c.Peso) / 1000.0
}

func TestandoCarro() {
	car := Car{
		Peso:  1500,
		Marca: "Toyota",
	}

	res, err := json.Marshal(car)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))
}
