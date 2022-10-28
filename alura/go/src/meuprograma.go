package main

import (
	"fmt"
)

func main() {
	nome := "Alexandre"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa esta na versão", versao)

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir os Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	// fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)
	fmt.Println("O endereço da minha variavel comando é", &comando)
	fmt.Println("O Comando escolhido foi", comando)

	// Scan pede o endereço da variavel por isso passamos assim &comando

	if comando == 1 {

	} else if comando == 2 {

	}

}
