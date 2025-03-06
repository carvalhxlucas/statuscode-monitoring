package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()

	for {
		exibeMenu()
		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 3:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando não encontrado!")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {

	fmt.Println("Seja bem-vindo ao sistema de monitoramento")

	var nome string
	fmt.Print("Digite o seu nome: ")
	fmt.Scan(&nome)

	fmt.Println("Olá, sr.", nome, "\n---------------")

}

func exibeMenu() {
	fmt.Print("1- Iniciar monitoramento\n" +
		"2- Exibir Logs\n" +
		"3- Sair do programa\n" +
		"Escolha uma opção: ")

}

func lerComando() int {
	var comandoInserido int
	fmt.Scan(&comandoInserido)

	return comandoInserido
}

func iniciarMonitoramento() {
	var urlSite string
	fmt.Print("Insira a URL do site: ")
	fmt.Scan(&urlSite)
	fmt.Println("Iniciando monitoramento ao site...")

	resp, _ := http.Get(urlSite)

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", urlSite, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site: ", urlSite, "está com problemas! StatusCode: ", resp)
	}
}
