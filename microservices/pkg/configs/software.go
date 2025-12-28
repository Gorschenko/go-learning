package configs

type SoftwareConfig struct {
	Api    ApiSoftwareConfig    `json:"api"`
	Logger LoggerSoftwareConfig `json:"logger"`
}

type LoggerSoftwareConfig struct {
	Level     string `json:"level"`
	AddSource bool   `json:"add_source"`
}

type ApiSoftwareConfig struct {
	TimeoutSec int  `json:"timeout_sec"`
	Debug      bool `json:"debug"`
}
