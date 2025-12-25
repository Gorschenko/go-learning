package configs

type SoftwareConfig struct {
	Api ApiSoftwareConfig `json:"api"`
}

type ApiSoftwareConfig struct {
	TimeoutSec int  `json:"timeout_sec"`
	Debug      bool `json:"debug"`
}
