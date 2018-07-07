package main

import (
	"fmt"
)

func main() {
	var numero int
	fmt.Print("Digite um número: ")
	fmt.Scanf("%d", &numero)

	// Aaaaa... pegadinha do malandro!!! não tem WHILE no Golang, e sim FOR, segura aí:
	i := 0
	for i <= 10 {
		fmt.Printf("%d x %d = %d \n", numero, i, numero*i)
		i++
	}
}
