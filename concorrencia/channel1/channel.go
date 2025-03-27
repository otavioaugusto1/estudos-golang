package main

import (
	"fmt"
)

func main(){
	ch := make(chan int, 1) // criando um channel de inteiros e com o buffer 1
	ch <- 1 //enviando dados para o canal
	<-ch
	ch <- 2
	fmt.Println(<-ch)

}