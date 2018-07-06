package main

import (
	"fmt"
)

func main() {

	// Declaração de Constante
	const minhaConstante string = "Curso Golang - Declarando Constantes"
	fmt.Println(minhaConstante)

	// Podemos também ter mais de uma constante na declaração. Podemos ainda usar inferência veja:
	const (
		primeiroNome = "Declarando mais de uma constante"
		segundoNome  = "e usando Inferência de tipos"
	)

	fmt.Println(primeiroNome + " " + segundoNome)
}
