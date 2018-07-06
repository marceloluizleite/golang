package main

import (
	"fmt"
)

func main() {
	fmt.Println("Inteiros sem sinal")

	var u1 byte = 255 //uint8
	fmt.Println(u1)

	var u2 uint16 = 33
	fmt.Println(u2)

	var u3 uint32 = 44
	fmt.Println(u3)

	var u4 uint64 = 55
	fmt.Println(u4)

	fmt.Println("------------------")
	fmt.Println("Inteiros com sinal")
	fmt.Println("------------------")

	var i1 int8 = 127
	fmt.Println(i1)

	var i2 int16 = 1000
	fmt.Println(i2)

	var i3 int32 = 1001 // ou posso usar a keyword: rune - Ao invés de int32
	fmt.Println(i3)

	var i4 int64 = 1002
	fmt.Println(i4)

	var t1 uint = 1010
	fmt.Println(t1)

	fmt.Println("------------------")
	fmt.Println("Pontos flutuantes")
	fmt.Println("------------------")

	var f1 float32 = 1
	fmt.Println(f1)

	var f2 float64 = 2
	fmt.Println(f2)

	fmt.Println("------------------")
	fmt.Println("Complexos")
	fmt.Println("------------------")

	var f3 complex64 = 3
	fmt.Println(f3) // a saída será: (3+0i)

	var f4 complex128 = 4
	fmt.Println(f4) // a saída será: (4+0i) - Util para notação científica.

	fmt.Println("------------------")
	fmt.Println("Strings")
	fmt.Println("------------------")

	var s1 string = "Treinamento Golang"
	var s2 string = "Documento o Processo no \ngit"
	fmt.Println(s1)
	fmt.Println(s2)

	var s3 string = `Treinamento Golang
	TreinaWeb`
	fmt.Println(s3)

	fmt.Println("------------------")
	fmt.Println("Booleanos")
	fmt.Println("------------------")

	var b1 bool = true
	fmt.Println(b1)

	fmt.Println("------------------")
	fmt.Println("Multiplas Declarações (variáveis)")
	fmt.Println("------------------")

	var nome, sobrenome string = "Anakin", "Skywalker"
	fmt.Println(nome + " " + sobrenome)

	// Ou ainda:

	var (
		filme  string = "Star Wars"
		titulo string = "Os últimos Jedi"
		ano    int    = 2017
	)
	fmt.Println(filme + " " + titulo)
	fmt.Println(ano)
}
