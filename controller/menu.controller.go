package controller

import (
	"financer/model"
	"fmt"
	"os"
	"time"
)

func MenuController() {
	var choice string

	fmt.Println("")
	fmt.Println("1. VISUALIZAR GASTOS \n2. INSERIR GASTOS \n3. CRIAR NOVO ARQUIVO \n4. CONFIG\n5. SAIR\n")

	_, err := fmt.Scan(&choice)

	if err != nil {
		fmt.Println("Tente novamente.")
		MenuController()
	} else {
		switch choice {
		case "1":
			{
				fmt.Print("Digite nome do arquivo: ")
				var arquivoEscolhido string
				_, err := fmt.Scan(&arquivoEscolhido)
				if err != nil {
					fmt.Println("ErrorLerEntradaMes: " + err.Error())
				}
				model.VisualizarGastos(arquivoEscolhido)
				MenuController()
			}
		case "2":
			{
				var gasto, prioridade string
				var valor float64
				var arquivo string
				fmt.Print("\nDigite o nome do arquivo que deseja inserir: ")
				_, err := fmt.Scan(&arquivo)
				if err != nil {
					fmt.Println("\nErro ao capturar a resposta do nome do arquivo: " + err.Error())
				} else {
					fmt.Print("\nDigite o nome do Gasto: ")
					_, err := fmt.Scan(&gasto)
					if err != nil {
						fmt.Println("Erro ao capturar a resposta do gasto: " + err.Error())
					} else {
						fmt.Print("\nDigite o valor do Gasto: ")
						_, err := fmt.Scan(&valor)
						if err != nil {
							fmt.Println("Erro ao capturar a resposta do valor: " + err.Error())
						} else {
							fmt.Print("\nDigite a prioridade: ")
							_, err := fmt.Scan(&prioridade)
							if err != nil {
								fmt.Println("Erro ao capturar a prioridade: " + err.Error())
							}
							model.InserirGastos(valor, arquivo, gasto, prioridade)
							MenuController()
						}
					}
				}
			}
		case "3":
			{
				var nomeArquivo string
				fmt.Print("Digite o nome do novo arquivo: ")
				fmt.Scan(&nomeArquivo)
				model.CriarNovo(nomeArquivo)
				MenuController()
			}
		case "4":
			{

			}
		case "5":
			{
				go func() {
					var count = 4
					for count > 0 {
						fmt.Printf("SAINDO EM %d ...\n", (count - 1))
						count -= 1
						time.Sleep(time.Second)
					}
				}()
				time.Sleep(time.Second * 3)
				os.Exit(0)
			}
		default:
			{
				fmt.Println("Opção inválida.")
				time.Sleep(time.Second * 1)
				MenuController()
			}
		}
	}
}
