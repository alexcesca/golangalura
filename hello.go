package main

import (
	"fmt"
	"net/http"
	"os"
)

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
	versao := 1.1

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
	site := "https://www.alura.com.br"
	retorno, err := http.Get(site)
	if err != nil {
		fmt.Printf("Site %v OffLine. Status code %v", site, retorno.Status)
	}
	if retorno.StatusCode == 200 {
		fmt.Printf("Site %v Online ", site)
	} else {
		fmt.Printf("Site %v OffLine. Status code %v", site, retorno.Status)
	}
}
