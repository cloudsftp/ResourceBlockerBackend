package server

import (
	"encoding/json"
	"os"
)

type Resource struct {
	Name string `json:"name"`
	Min  int    `json:"min"`
	Max  int    `json:"max"`
}

type Config struct {
	Port      int        `json:"port"`
	Resources []Resource `json:"resources"`
}

func GetConfig(configFilePath string) Config {
	configFile, err := os.Open(configFilePath)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	configDecoder := json.NewDecoder(configFile)
	config := &Config{}
	err = configDecoder.Decode(config)
	if err != nil {
		panic(err)
	}

	return *config
}
