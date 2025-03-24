package main

import "fmt"

func f1(){
	fmt.Println("F1")
}

func f2(p1 string){
	fmt.Println(p1)

}

func f3(p2 string) string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string){
	return "Retorno 1", "Retorno 2"
}

func main(){
	f1()
	f2("Aoba")

	r1 := f3("aoba")
	fmt.Println(r1)

	r2 := f4("valo1", "valor2")
	fmt.Println(r2)
}