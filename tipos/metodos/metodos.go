package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome string
	sobrenome string
}
// método com receiver (recebedor) 
func (p pessoa) getNomeCompleto() string{
	return p.nome + " " + p.sobrenome
}
// usando ponteiro pois preciso alterar, se fosse somente leitura não iria precisar
func (p *pessoa) setNomeCompleto(nomeCompleto string){
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]

}

func main(){
	p1 := pessoa{"Otavio", "Nogueira"}
	fmt.Println(p1)
	fmt.Println(p1.getNomeCompleto())
	p1.setNomeCompleto("Gonçalves Nogueira")
	fmt.Println(p1.getNomeCompleto())
}

