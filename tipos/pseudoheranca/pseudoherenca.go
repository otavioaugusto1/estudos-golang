package main

import "fmt"

type carro struct {
	nome string
	velocidadeAtual int
}
type ferrari struct {
	carro // campos anonimos
	turboLigado bool	
}

func main(){	
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true
	fmt.Println(f)
	fmt.Printf("A ferrari com nome %v está com turbo ligado? %v", f.nome, f.turboLigado)

}
