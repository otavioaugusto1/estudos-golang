package main

import "fmt"

func main(){
	//arrays são homogêneos (mesmo tipo) e estáticos (fixo)
	var notas [3] float64
	fmt.Println(notas)
	notas[0], notas[1], notas[2] = 7.0,6.3,5.5
	fmt.Println(notas)
	
	total := 0.0
	for i := 0; i < len(notas); i++{
		total += notas[i]
	}
	media := total / float64(len(notas))
	fmt.Println(media)

	// usando range no array
	numeros := [...]int{1,2,3,4,5} //compilador que conta a qtd de posições do array

	for i, numero := range numeros{
		fmt.Printf("%d) %d \n", i, numero)
	}

}