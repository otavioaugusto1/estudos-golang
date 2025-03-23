package main

import "fmt"

func main(){
	a1 := [3]int{1,2,3} //array
	s1 := []int{1,2,3} // slice
	fmt.Println(a1, s1)

	a2 := [5]int{1,2,3,4,5} 
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	// mostrando que 2 slices apontam para o mesmo array
	slice1 := make([]int, 10, 20)
	slice2 := append(slice1, 1, 2, 3)
	fmt.Println(slice1, slice2)

	slice1[0] = 7
	fmt.Println(slice1, slice2)

}