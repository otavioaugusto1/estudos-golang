package main

import (
	"fmt"
)

func p1(n int){
	n++
	fmt.Println("Print de n dentro da função p1",n)
}

func p2(n *int){
	*n++
	fmt.Println("Print de n dentro da função p2",n)

}

func main(){
	valor := 3
	fmt.Println("Valor da variavel dentro da main", valor)
	p1(valor)
	fmt.Println("Valor da variavel dentro da main após chamar a função sem ponteiro", valor)

	p2(&valor)
	fmt.Println("Valor da variavel dentro da main após chamar a função com ponteiro", valor)


}