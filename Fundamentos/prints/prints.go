package main

import "fmt"

func main() {
	fmt.Print("mesma")
	fmt.Print(" linha.")

	fmt.Println(" Linha")
	fmt.Println("nova")

	x := 3.141516
	//fmt.Println("O valor de x é" + x)
	xs := fmt.Sprint(x)
	fmt.Println("O valor de xs é " + xs)
	fmt.Println("O valor de x é",  x)
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n %d %f %t %s", a, b, c, d)

}