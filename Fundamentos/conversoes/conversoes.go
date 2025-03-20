package main

import (
	"fmt"
	"strconv"
)

func main(){
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	// Cuidado, da forma abaixo está pegando o 123 da tabela ascii
	fmt.Println("Teste " + string(123))

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// String para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)
	
}