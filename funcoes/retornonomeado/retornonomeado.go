package main

import (
	"fmt"
)

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo, pois as variáveis de retorno já receberam seus valores
}


func main(){
	v1, v2 := trocar(1, 2)
	fmt.Println(v1, v2)
}