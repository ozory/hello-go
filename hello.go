package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	for {
		fmt.Println("1 - Iniciar Monitoramento")
		fmt.Println("2 - Exibir Logs")
		fmt.Println("0 - Sair do Programa")

		exibeMenu()
	}
}

func lerComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

// Exibe o menu
func exibeMenu() {
	switch lerComando() {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Nenhuma opção")
	}
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando")

	sites := []string{
		"https://teovest.com.br",
		"https://priceclub.com.br",
	}

	// For interativo i := 0; i < len(sites); i++
	// index, site := range sites
	for _, site := range sites {
		response, error := http.Get(site)
		if error != nil {
			fmt.Println("Houve um erro", error.Error())
			os.Exit(-1)
		}

		if response.StatusCode == 200 {
			fmt.Println("Site", site, "funcionando")
		} else {
			fmt.Println("Site não encontrado ou com problemas", response.StatusCode)
			os.Exit(-1)
		}
	}

}
