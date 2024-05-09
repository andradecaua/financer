package model

import (
	"encoding/csv"
	"financer/services"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Vizualizar gastos do mês atual
func VisualizarGastos(fileName string) {

	var gastoTotal float64
	var resultado string
	var renda float64 = 0

	config := services.AbrirConfig()

	renda = config.Salario
	// Abrir csv
	file, err := os.OpenFile(fmt.Sprintf("meses/%s.csv", fileName), os.O_RDWR, os.ModeDir)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo: " + err.Error())
	} else {
		// Ler arquivo CSV
		dados, err := os.ReadFile(file.Name())
		if err != nil {
			fmt.Println("Erro ao ler o arquivo: " + err.Error())
		}
		csvReader := csv.NewReader(strings.NewReader(string(dados)))
		csvReader.Comma = rune(';')
		dadosCSV, err := csvReader.ReadAll()
		if err != nil {
			fmt.Println("Erro ao converter o csv: " + err.Error())
		}
		fmt.Println("-----------------------------------------------")
		for _, rows := range dadosCSV {
			row := fmt.Sprintf("%s\t%s\t%s ", rows[0], rows[1], rows[2])
			convertedValue, err := strconv.ParseFloat(rows[1], 64)
			if err != nil {
				if rows[1] != "VALOR" {
					fmt.Println("Erro ao converter o valor: " + err.Error())
				}
			}
			gastoTotal += convertedValue
			fmt.Println("\t" + row)
		}

		resultado = func() string {
			if renda > gastoTotal {
				return "POSITIVO"
			} else {
				return "NEGATIVO"
			}
		}()

		fmt.Printf("\n\tGasto total: R$%.2f\n", gastoTotal)
		fmt.Printf("\n\tSaldo: R$%2.f\n", (renda - gastoTotal))
		fmt.Println("\n\tResultado: " + resultado)
		fmt.Println("-----------------------------------------------")
	}
}
