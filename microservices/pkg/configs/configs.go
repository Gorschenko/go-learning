package configs

import (
	"encoding/json"
	"os"
)

type Config struct {
	Database DatabaseConfig `json:"database"`
	Services ServicesConfig `json:"services"`
}

type ServicesConfig struct {
	Auth   ServiceConfig `json:"auth"`
	Orders ServiceConfig `json:"orders"`
}

type ServiceConfig struct {
	Port int    `json:"port"`
	Host string `json:"host"`
}

type DatabaseConfig struct {
	ServiceConfig
	Database    string `json:"database"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Automigrate bool   `json:"automigrate"`
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
