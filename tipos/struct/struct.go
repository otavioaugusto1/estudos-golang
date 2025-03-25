package main

import "fmt"

type produto struct {
	nome string
	preco float64
	desconto float64
}
// método com receiver (recebedor) 
func (p produto) precoComDesconto() float64{
	return p.preco * (1 - p.desconto)
}

func main(){
	var produto1 produto
	produto1 = produto{
		nome: "Lapis",
		preco: 1.79,
		desconto: 0.05,
	}
	fmt.Println(produto1)
	fmt.Printf("O produto %v com desconto de %v fica = %v", produto1.nome, produto1.desconto, produto1.precoComDesconto())
	produto2 := produto{"notebook", 3.500, 0.10}
	fmt.Println(" ")

	fmt.Println(produto2)
	
}
