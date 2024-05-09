package model

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rendasExtras struct {
	Nome  string
	Valor float64
}

type config struct {
	Salario      float64        `json:"Salario"`
	RendasExtras []rendasExtras `json:"rendaExtra"`
}

// Vizualizar gastos do mês atual
func VisualizarGastos(fileName string) {

	var gastoTotal float64
	var resultado string
	var renda float64 = 0
	var config config

	configFile, errJson := os.OpenFile("config.json", os.O_RDONLY, os.ModeDir)

	if errJson != nil {
		fmt.Println("Erro ao abrir as config: " + errJson.Error())
	}

	dadosConfig, errConfig := os.ReadFile(configFile.Name())

	if errConfig != nil {
		fmt.Println("Erro ao ler as config: " + errConfig.Error())
	}

	errUnmarshal := json.Unmarshal(dadosConfig, &config)
	if errUnmarshal != nil {
		fmt.Println("Erro ao converter a config: " + errUnmarshal.Error())
	}

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
