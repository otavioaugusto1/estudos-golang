package main

import "fmt"

func soma(a int, b int) int {
	return a + b 
}

func imprimir(valor int) {
	fmt.Println(valor)
}

func main() {
	resultado := soma(2, 2)
	imprimir(resultado)
}