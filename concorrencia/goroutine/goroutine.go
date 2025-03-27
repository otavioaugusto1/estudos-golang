package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i:= 0; i < qtde; i++{
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (interação %v) \n", pessoa, texto, i+1)
	}
}
//Goroutine: funções que são executadas de forma independente
func main(){
	//fale("Maria", "Estou falando", 3)
	//fale("João", "Só falo depois que você terminar", 1)
	
	//fmt.Println

	//go fale("Maria", "Ei", 100)
	//go fale("Joao", "OPA", 100)

	go fale("Maria", "Opa", 10)
	fale("Joaõ", "Ei", 5)

}