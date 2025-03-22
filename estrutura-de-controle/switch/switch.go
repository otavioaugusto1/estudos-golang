package main

import (
"fmt"
)

func notaParaConceito(n float64) string{
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough
case 9:
	return "A"
case 7,8:
	return "B"
default:
	return "nota inválida"
	}
}
func main() {
	fmt.Println(notaParaConceito(5.0))
}