package main

import (
	"fmt"
)

// funções anônimas

var soma = func(a, b int) int {
	return a + b
}

func main(){
	fmt.Println(soma(1,2))

	sub := func(c, d int) int {
		return c - d
	}
	fmt.Println(sub(2,1))
}