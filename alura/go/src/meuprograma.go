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
	// fmt.Println("O endereço da minha variavel comando é", &comando)
	// fmt.Println("O Comando escolhido foi", comando)

	// Scan pede o endereço da variavel por isso passamos assim &comando

	// if comando == 1 {
	// 	fmt.Println("Monitorando...")

	// } else if comando == 2 {
	// 	fmt.Println("Exibindo Logs...")

	// } else if comando == 0 {
	// 	fmt.Println("Saindo do Programa")

	// } else {
	// 	fmt.Println("Não conheço este comando")
	// }

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
		break
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do Programa")
	default:
		fmt.Println("Não conheço este comando")

	}

}
