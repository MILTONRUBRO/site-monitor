package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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

	//urls := []string{"https://www.alura.com.br", "https://random-status-code.herokuapp.com/", "https://www.google.com/"}

	urls := lerUrlsDeArquivo()

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
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	fmt.Println(string(arquivo))
}

func testarUrls(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site", url, "foi carregado com sucesso")
		registrarLog(url, true)
	} else {
		fmt.Println("Erro ao carregar o site", url, "Codigo de erro ", resp.StatusCode)
		registrarLog(url, false)
	}

	fmt.Println()
}

func lerUrlsDeArquivo() []string {

	var urls []string

	//arquivo, err := ioutil.ReadFile("urls.txt")

	arquivo, err := os.Open("urls.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {

		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		urls = append(urls, linha)

		if err != nil {
			fmt.Println("Ocorreu um erro na leitura do arquivo", err)
		}

		if err == io.EOF {
			break
		}
	}

	//fechando arquivo
	arquivo.Close()

	return urls
}

func registrarLog(url string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro")
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + url + "- Online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()

}
