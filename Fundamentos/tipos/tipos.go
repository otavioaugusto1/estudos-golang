package main

import (
	"fmt"
	"reflect"
)

func main() {
	//só sinais positivos: uint8, uint16, uint32, uint64
	var b byte = 255
	fmt.Println("O byte é ", reflect.TypeOf(b))
	//Positivos e negativos int8 int16 int32 int64

	//números reais (float32, float64)
	var x float64 = 49.99
	fmt.Println("O tipo de x  é ", reflect.TypeOf(x))
	fmt.Println("O valor é ",b)

	//tipo booleano
	var boo bool = true
	fmt.Println("O tipo de x  é ", reflect.TypeOf(boo))
	fmt.Println("O valor é ",boo)


}