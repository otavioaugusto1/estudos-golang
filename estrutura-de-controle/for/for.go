package main

import (
"fmt"
)



func main() {
	i := 1
	for i <= 10 {
		fmt.Println("Valor de i ", i)
		i++
	}
	for i := 0; i <= 20; i +=2{
		fmt.Printf("%v", i)
	}
}