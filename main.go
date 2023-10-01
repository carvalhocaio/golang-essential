package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Hailey"
	var versao float32 = 1.2
	fmt.Println("Olá,", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := []string{
		"https://random-status-code.onrender.com",
		"https://alura.com.br",
		"https://caelum.com.br",
	}

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			fmt.Println("Testando site:", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println(">>", site, "foi carregado com sucesso!")
	} else {
		fmt.Println(">>", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
