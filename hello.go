package main

import (
	"fmt"
	"os"
)

type Carro struct {
	Placa string
	Modelo string
}

var (
	listCarros = []Carro{}
)
func main() {
	
	for {
		exibirMenu()

		comando := lerComando()

		switch comando {
		case 1:
			inserirCarros()
		case 2:
			listarCarros()
		case 3:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Printf("Não conheço este comando")
			os.Exit(-1)
		}
		
	}

}

func exibirMenu() {
	fmt.Println("-------------ESTACIONAMENTO----------------")
	fmt.Println("1 - Inserir carro")
	fmt.Println("2 - Listar carros")
	fmt.Println("3 - Saindo do programa")
}

func lerComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func inserirCarros() {
	var modelo, placa string
	fmt.Println("Digite o modelo")
	fmt.Scan(&modelo)
	fmt.Println("Digite a placa")
	fmt.Scan(&placa)

	carro := Carro{
		 Placa: placa,
		 Modelo: modelo,
	}

	listCarros = append(listCarros, carro)

	var result = fmt.Sprintf("O carro adicionado foi %s com a placa %s", carro.Modelo, carro.Placa)

	fmt.Println(result)

}

func listarCarros() {
	if len(listCarros) == 0{
		fmt.Println("Não há carros no estacionamento")
		return
	}
	for _, carro := range listCarros {
		fmt.Println(carro.Modelo)
	}
}


