package main

import (
	"fmt"
)

func main() {
	var mes int
	fmt.Println("Digite o número do mês: ")
	fmt.Scanf("%d", &mes)

	switch mes {
	case 1:
		fmt.Println("Este mês é Janeiro.")
	case 2:
		fmt.Println("Este mês é Fevereiro.")
	case 3:
		fmt.Println("Este mês é Março.")
	case 4:
		fmt.Println("Este mês é Abril.")
	case 5:
		fmt.Println("Este mês é Maio.")
	case 6:
		fmt.Println("Este mês é Junho.")
	case 7:
		fmt.Println("Este mês é Julho.")
	case 8:
		fmt.Println("Este mês é Agosto.")
	case 9:
		fmt.Println("Este mês é Setembro.")
	case 10:
		fmt.Println("Este mês é Outubro.")
	case 11:
		fmt.Println("Este mês é Novembro.")
	case 12:
		fmt.Println("Este mês é Dezembro.")
	default:
		fmt.Println("Mês inválido!")
	}
	switch mes {
	case 1, 2, 3:
		fmt.Println("Primeiro trimestre.")
	case 4, 5, 6:
		fmt.Println("Segundo trimestre.")
	case 7, 8, 9:
		fmt.Println("Terceiro trimestre.")
	case 10, 11, 12:
		fmt.Println("Quarto trimestre.")
	}
}
