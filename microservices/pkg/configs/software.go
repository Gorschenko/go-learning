package configs

type SoftwareConfig struct {
	Api    ApiSoftwareConfig    `json:"api"`
	Logger LoggerSoftwareConfig `json:"logger"`
}

type LoggerSoftwareConfig struct {
	Level string `json:"level"`
}

type ApiSoftwareConfig struct {
	TimeoutSec int  `json:"timeout_sec"`
	Debug      bool `json:"debug"`
}
