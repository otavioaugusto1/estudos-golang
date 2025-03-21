package main

import "fmt"

func main(){
	i := 1

	var p *int = nil

	p = &i

	*p++
	i++
	fmt.Println(p, *p, i, &i) // endereço de memória, valor e valor
}