package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415 // declaração manual
	var raio = 3.2 //tipo (float64) inferido automaticamente

	//declaração reduzida
	area := PI * math.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
	)
	fmt.Println(a, b)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "opa"

	fmt.Println(g, h, i)

	
}