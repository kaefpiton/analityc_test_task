package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Router struct {
		Port string `json:"port"`
	}

	Repository struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		DBName   string `json:"db_name"`
		SSLMode  string `json:"ssl_mode"`
	}

	Logger struct {
		Lvl string `json:"lvl"`
	}
}

func LoadConfiguration(path string) (*Config, error) {
	var config *Config

	configFile, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
