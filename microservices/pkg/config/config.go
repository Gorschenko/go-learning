package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Services ServicesConfig `json:"services"`
}

type ServicesConfig struct {
	Auth   AuthServiceConfig   `json:"auth"`
	Orders OrdersServiceConfig `json:"orders"`
}

type AuthServiceConfig struct {
	Port int    `json:"port"`
	Host string `json:"host"`
}

type OrdersServiceConfig struct {
	Port int    `json:"port"`
	Host string `json:"host"`
}

func LoadConfig(filePath string) (*Config, error) {
	jsonConfig, err := os.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	var config Config

	err = json.Unmarshal(jsonConfig, &config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}
