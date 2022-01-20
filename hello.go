package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibeintroducao()
	for {
		exibeMenu()

		comando, err := leComando()
		if err != nil {
			fmt.Println("Opção invalida ", comando)
		}

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs..")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando não conhecido")
			os.Exit(-1)
		}
	}
}

func exibeintroducao() {
	nome := "Alex"
	versao := 1.2

	fmt.Println("Ola, Sr. ", nome)
	fmt.Println("Este programa esta na versão ", versao)
}

func leComando() (int, error) {
	var comando int
	_, err := fmt.Scan(&comando)

	if err != nil {
		fmt.Println("Opção invalida ", comando)
		return 0, err
	}
	return comando, nil
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}
func iniciarMonitoramento() {
	fmt.Println("Monitorando ...")
	// função de slice informando os parametros.
	//sites := []string{"https://random-status-code.herokuapp.com", "https://www.alura.com.br", "https://auroraalimentos.com.br/"}
	sites := leSitesDoArquivo()
	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		// função de time para parar 5 segundos.
		time.Sleep(delay * time.Second)
	}

	fmt.Println("")
}
func testaSite(site string) {
	retorno, err := http.Get(site)
	if err != nil {
		fmt.Printf("Site %v OffLine. Status code %v \n", site, retorno.Status)
	}
	if retorno.StatusCode == 200 {
		fmt.Printf("Site %v Online \n", site)
	} else {
		fmt.Printf("Site %v OffLine. Status code %v\n", site, retorno.Status)
	}
}
func leSitesDoArquivo() []string {
	// Abrindo arquivo go
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		arquivo.Close()
		log.Fatal(err.Error())
	}
	//a biblioteca ioutil restorna em um array de bites.
	//arquivo, err := ioutil.ReadFile("sites.txt")
	leitor := bufio.NewReader(arquivo)

	var sites []string
	for {
		linha, err := leitor.ReadString('\n')
		// a função TrimSpace remove os \n e espaços.
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		//fmt.Println(linha)
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}
