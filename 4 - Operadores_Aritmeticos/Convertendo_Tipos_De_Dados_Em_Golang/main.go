package main

import (
	"fmt"
	"strconv"
)

func main() {
	var texto string
	fmt.Print("Digite um número: ")
	fmt.Scanf("%s", &texto)
	var numero int
	numero, _ = strconv.Atoi(texto)
	/* Poderiamos ainda usar o método a seguir:

	numero, _ := strconv.ParseInt(texto, 10, 32) // --> Onde 10 é o formato e 32(bits) a precisão.
	*/
	fmt.Println(numero)

	// Cast em Golang:
	var conversao = float64(numero)
	fmt.Println(conversao)

	/*
		Em golang, o cast deve ser explicito e aí entra a inferência, pois poderiamos fazer:

		var conversao float64 = float64(numero)

		--> Neste caso Golang presa pela simplicidade então entra em cena a inferência entendendo que não é necessário repetir o float64.

	*/
}
