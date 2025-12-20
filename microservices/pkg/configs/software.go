package configs

type SoftwareConfig struct {
	Api ApiSoftwareConfig
}

type ApiSoftwareConfig struct {
	TimeoutSec int `json:"timeout_sec"`
}
