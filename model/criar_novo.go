package model

import (
	"fmt"
	"os"
)

// Criar novo arquivo criar um novo arquivo csv.
func CriarNovo(nomeArquivo string) {
	file, err := os.Create(fmt.Sprintf("meses/%s.csv", nomeArquivo))
	if err != nil {
		fmt.Println("Erro ao criar arquivo: " + err.Error())
	} else {
		_, err := file.Write([]byte("GASTOS;VALOR;PRIORIDADE"))
		if err != nil {
			fmt.Println("Erro ao gerar dados default: " + err.Error())
		} else {
			fmt.Println("Arquivo gerado com sucesso!")
		}
	}
}
