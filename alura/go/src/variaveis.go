package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Alexandre"
	idade := 37
	versao := 1.1
	fmt.Println("Olá, sr.", nome, "sua idade é ", idade)
	fmt.Println("Este programa esta na versão", versao)
	fmt.Println("O tipo da variavel nome é ", reflect.TypeOf(nome))
	fmt.Println("O tipo da variavel idade é ", reflect.TypeOf(idade))
	fmt.Println("O tipo da variavel idade é ", reflect.TypeOf(versao))

}

// Se eu não declarar a variavel o go assume o que ela é.
// operador de atribuição de variaveis curto := não precisa escrever a palavra var antes.
