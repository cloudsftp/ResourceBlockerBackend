package server

import (
	"encoding/json"
	"os"

	"github.com/cloudsftp/ResourceBlockerBackend/resource"
)

type Config struct {
	Port      int                           `json:"port"`
	Resources map[string]*resource.Resource `json:"resources"`
}

func GetConfig(configFilePath string) *Config {
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

	return config
}
