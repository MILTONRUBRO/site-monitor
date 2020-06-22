package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 10

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

	urls := []string{"https://www.alura.com.br", "https://random-status-code.herokuapp.com/", "https://www.google.com/"}

	for i := 0; i < monitoramentos; i++ {

		for i, url := range urls {
			fmt.Println("Testando o site", i, ":", url)
			testarUrls(url)
		}

		fmt.Println("--------------------------------------------------------------------")
		time.Sleep(delay * time.Second)
	}

}

func exibirLogs() {
	// to do
	fmt.Println("Exibindo Logs")

}

func testarUrls(url string) {
	resp, _ := http.Get(url)

	if resp.StatusCode == 200 {
		fmt.Println("Site", url, "foi carregado com sucesso")
	} else {
		fmt.Println("Erro ao carregar o site", url, "Codigo de erro ", resp.StatusCode)
	}

	fmt.Println()
}
