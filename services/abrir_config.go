package services

import (
	"encoding/json"
	"fmt"
	"os"
)

type rendasExtras struct {
	Nome  string
	Valor float64
}

type Config struct {
	Salario      float64        `json:"Salario"`
	RendasExtras []rendasExtras `json:"rendaExtra"`
}

// Abrir Config abre as configurações e retorna os dados em uma struct
func AbrirConfig() Config {

	var config Config
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
	return config
}
