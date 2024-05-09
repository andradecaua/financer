package model

import (
	"fmt"
	"os"
)

func InserirGastos(valor float64, arquivo, gasto, prioridade string) {

	dadosNovos := []byte(fmt.Sprintf("\n%s;%f;%s", gasto, valor, prioridade))
	file, err := os.OpenFile(fmt.Sprintf("meses/%s.csv", arquivo), os.O_RDWR, os.ModeDir)

	if err != nil {
		fmt.Println("Erro Abrir arquivo: " + err.Error())
	} else {
		//Ler csv atual
		dados, err := os.ReadFile(file.Name())
		if err != nil {
			fmt.Println("Erro ao ler o arquivo: " + err.Error())
		} else {
			dados = append(dados, dadosNovos...)
			file.Write(dados)
			file.Close()
			fmt.Println("\nGastos Inseridos com sucesso!")
		}
	}
}
