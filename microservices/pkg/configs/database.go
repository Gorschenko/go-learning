package configs

type DatabaseConfig struct {
	ServiceConfig
	Database    string `json:"database"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Automigrate bool   `json:"automigrate"`
}
