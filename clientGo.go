package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	for i := 0; i < 1000; i++ {
		var opcao = ""
		var nome = ""
		var idade = ""
		var ocupacao = ""
		//var cpf = ""
		//var num1 = 0
		//var num2 = 0
		//var operador = ""

		//fmt.Print("\033[H\033[2J")
		fmt.Printf("1 para adicionar usuário\n")
		fmt.Printf("2 para visualizar usuário\n")
		fmt.Printf("3 para editar usuário\n")
		fmt.Printf("4 para deletar usuário\n")
		fmt.Printf("5 para verificar cpf\n")
		fmt.Printf("6 para calcular operação\n")
		fmt.Printf("s para sair\n")
		fmt.Scanln(&opcao)

		if opcao == "S" || opcao == "s" {
			break
		}
		if opcao != "1" && opcao != "2" && opcao != "3" && opcao != "4" && opcao != "5" && opcao != "6" && opcao != "s" && opcao != "S" {
			fmt.Printf("Opção inválida, tente novamente\n")
			fmt.Scanln(&opcao)
		}
		if opcao == "1" {

			fmt.Printf("Nome\n")
			fmt.Scanln(&nome)

			fmt.Printf("Idade\n")
			fmt.Scanln(&idade)

			fmt.Printf("Ocupação\n")
			fmt.Scanln(&ocupacao)

		}
		if opcao == "2" {
			fmt.Printf("Nome\n")
			fmt.Scanln(&nome)

			response, err := http.Get("http://localhost:5000/user/" + nome)
			if err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", err)
			} else {
				data, _ := ioutil.ReadAll(response.Body)
				fmt.Println(string(data))
			}
			fmt.Scanln(&nome)
		}
		if opcao == "3" {
			fmt.Printf("Nome\n")
			fmt.Scanln(&nome)
		}
		if opcao == "4" {
			fmt.Printf("Nome\n")
			fmt.Scanln(&nome)

			client := &http.Client{}

			req, err := http.NewRequest("DELETE", "http://localhost:5000/user/"+nome, nil)
			if err != nil {
				fmt.Println(err)
				return
			}

			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()
		}
	}
}
