package main

import "fmt"

func main(){
	aprovados := make(map[int]string)
	aprovados[12345567] = "maria"
	aprovados[123335567] = "Joao"
	aprovados[1235677567] = "Pedro"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF %d) \n", nome, cpf)
	}

	fmt.Println(aprovados[12345567])
	delete(aprovados, 12345567)
	fmt.Println(aprovados)

	// já definindo o valor do map na inicialização
	funcESalarios := map[string]float64 {
		"jose joão": 11335.45,
		"maria heeheh": 1000.30,
		"Pedro Raul": 1200.00,
	}

}