package configs

import (
	"encoding/json"
	"os"
)

type Config struct {
	Database     DatabaseConfig     `json:"database"`
	InitDatabase InitDatabaseConfig `json:"init_database"`
	Software     SoftwareConfig     `json:"software"`
	Services     ServicesConfig     `json:"services"`
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
