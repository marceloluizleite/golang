package main

import (
	"fmt"
)

func main() {
	var idade int
	fmt.Print("Digite sua idade: ")
	fmt.Scanf("%d", &idade)

	if idade >= 18 {
		fmt.Println("Você é maior de idade!")
	} else {
		fmt.Println("Você é menor de idade!")
	}

	var numero int
	fmt.Println("Digite um número: ")
	fmt.Scanf("%d", &numero)
	if numero < 0 {
		fmt.Println("Número negativo")
	} else if numero >= 0 && numero <= 9 {
		fmt.Println("Numero positivo e contém um digito")
	} else {
		fmt.Println("Número positivo, e contém mais de um digito")
	}
}
