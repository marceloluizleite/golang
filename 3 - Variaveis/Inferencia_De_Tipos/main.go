package main

import "fmt"

func main() {

	var nome = "Marcelo"
	fmt.Println(nome)

	/*
	 Ou seja, Golang entende que "Marcelo" é uma string mesmo sem que eu tenha declarado
	 seu tipo. Isso não é tipagem fraca, e sim inferência de tipo.

	*/
	// Short Initialization -- Aqui podemos declara a variável de forma direta sem o tipo e sem a palavra "VAR".
	nomeCompleto := "Curso de Golang"
	fmt.Println(nomeCompleto)

}
