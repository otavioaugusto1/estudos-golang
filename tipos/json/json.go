package main

import (
	"fmt"
	"encoding/json"
)

type produto struct {
	ID int `json:"id"`
	Nome string `json:"json"`
	Preco float64 `json:"preco"`
	Tags []string `json:"tags"`
}


func main(){	
	p1 := produto{1, "Notebook", 1.899, []string{"Promoção", "Eletrônico"}}
	p1Json,_ := json.Marshal(p1) 
	fmt.Println(p1)
	fmt.Println(p1Json)



}
