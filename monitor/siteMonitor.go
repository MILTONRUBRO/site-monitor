package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibirIntroducao()

	for {

		exibirMenu()

		//le opcao do usuario
		opcao := lerOpcao()

		switch opcao {
		case 1:
			iniciarMonitoramento()
		case 2:
			exibirLogs()
		case 0:
			sairPrograma()
		default:
			fmt.Println("Opção invalida")
		}

	}

}

func exibirIntroducao() {
	nome := "Robot"
	versao := 0.1

	fmt.Println("Bem  vindo Sr.", nome)
	fmt.Println("O sistema está na versão", versao)
}

func exibirMenu() {

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair ")
}

func lerOpcao() int {
	var opcao int
	fmt.Scan(&opcao)

	fmt.Println("A opção escolhida foi", opcao)

	return opcao
}

func sairPrograma() {
	fmt.Println("Saindo do programa...")
	os.Exit(0)
}

func iniciarMonitoramento() {
	// to do
	fmt.Println("iniciando monitoramento...")
	url := "https://www.alura.com.br"
	resp, _ := http.Get(url)
	//fmt.Println("Resposta", resp)

	if resp.StatusCode == 200 {
		fmt.Println("Site", url, "foi carregado com sucesso")
	} else {
		fmt.Println("Erro ao carregar o site", url)
	}
}

func exibirLogs() {
	// to do
	fmt.Println("Exibindo Logs")

}
