package main

import (
	"fmt"
	"log"
)

func main() {
	nome := "alex"
	versao := 1.1

	fmt.Println("Ola, Sr. ", nome)
	fmt.Println("Este programa esta na versão ", versao)

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

	var comando int
	_, err := fmt.Scan(&comando)

	if err != nil {
		fmt.Println("Opção invalida ", comando)
		log.Fatal(err.Error())
	}

	if comando == 1 {
		fmt.Println("Monitorando ...")

	} else if comando == 2 {
		fmt.Println("Exibindo logs..")

	} else if comando == 0 {
		fmt.Println("Saindo do programa...")
	} else {
		fmt.Println("Comando não conhecido")
	}

	switch comando {
	case 1:
		fmt.Println("Monitorando ...")
	case 2:
		fmt.Println("Exibindo logs..")
	case 0:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Comando não conhecido")
	}

}
