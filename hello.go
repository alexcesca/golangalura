package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
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
			imprimeLosg()
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
		registraLog(site, true, "")
	} else {

		registraLog(site, false, retorno.Status)
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
func registraLog(site string, status bool, erro string) {

	// open file da a opção de criar, escrever entre ourras.
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		arquivo.Close()
		log.Fatal(err.Error())

	}
	// para formatar a data tem que consultar a biblioteca do GO, por que ele trabalha com formado diferente, colocando numeros que representam os campos.
	linha := time.Now().Local().Format("02/01/2006 15:04:05") + " " + site + " - online: " + strconv.FormatBool(status)
	if erro != "" {
		linha = linha + " : " + erro
	}
	linha = linha + "\n"
	arquivo.WriteString(linha)
	arquivo.Close()
}
func imprimeLosg() {
	fmt.Println("Exibindo logs..")
	// na função de ioutil, já fecha o arquivo
	// tambem le todo o arquivo,
	//   faz o processo de abrir, ler o arquivo todo e fecha o arquivo.
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {

		fmt.Println(err.Error())
	}
	fmt.Println(string(arquivo))
}
