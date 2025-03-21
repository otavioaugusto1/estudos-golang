package main

import "fmt"

func compras(trab1 bool, trab2 bool) (bool, bool, bool){
	comprarTv50 := trab1 && trab2 //AND E
	comprarTv32 := trab1 != trab2 // OU Exclusivo
	comprarSorvete := trab1 || trab2 // OU

	return comprarTv50, comprarTv32, comprarSorvete
}

func main(){
	tv50, tv32, sorvete := compras(true, true)
	fmt.Println(tv50, tv32, sorvete)
	fmt.Printf("tv50 %t, tv32 %t, sorvete %t, saudável %t", tv50, tv32, sorvete, !sorvete)
}