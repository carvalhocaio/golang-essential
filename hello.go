package main

import "fmt"

func main() {
	nome := "Hailey"
	var versao float32 = 1.1
	fmt.Println("Hello,", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")

	var comando int
	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi", comando)

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo Logs...")
	} else if comando == 0 {
		fmt.Println("Saindo...")
	} else {
		fmt.Println("Não conheço esse comando")
	}
}
