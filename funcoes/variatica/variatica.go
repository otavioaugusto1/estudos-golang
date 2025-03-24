package main

import (
	"fmt"
)

// Função com a quantidade indefinida de parâmetro
func media (numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main(){
	fmt.Printf("A média é %.2f", media(5.5,7.4,5.9))


}