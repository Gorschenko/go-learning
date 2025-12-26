package configs

import (
	"encoding/json"
	"log/slog"
	"os"
)

type Config struct {
	Other        OtherConfig        `json:"other"`
	Database     DatabaseConfig     `json:"database"`
	InitDatabase InitDatabaseConfig `json:"init_database"`
	Software     SoftwareConfig     `json:"software"`
	Services     ServicesConfig     `json:"services"`
	Security     SecurityConfig     `json:"security"`
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

	slog.Debug(
		"Config is loaded",
		"Data",
		config,
	)

	return &config, nil
}
