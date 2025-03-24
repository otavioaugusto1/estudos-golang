package main

import (
	"fmt"
)

// passando uma função no parâmetro de outra função

func multiplica(a, b int) int {
	return a * b
} 

func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main(){
	fmt.Println(exec(multiplica,1, 2))


}