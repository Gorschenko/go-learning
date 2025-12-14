package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Databases DatabasesConfig `json:"databases"`
	Services  ServicesConfig  `json:"services"`
}

type DatabasesConfig struct {
	Pg PgDatabaseConfig `json:"pg"`
}

type ServicesConfig struct {
	Auth   ServiceConfig `json:"auth"`
	Orders ServiceConfig `json:"orders"`
}

type ServiceConfig struct {
	Port int    `json:"port"`
	Host string `json:"host"`
}

type PgDatabaseConfig struct {
	ServiceConfig
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
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
