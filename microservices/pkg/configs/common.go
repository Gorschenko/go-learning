package configs

type ServicesConfig struct {
	Auth   ServiceConfig `json:"auth"`
	Orders ServiceConfig `json:"orders"`
}

type ServiceConfig struct {
	Port     int    `json:"port"`
	Host     string `json:"host"`
	Protocol string `json:"protocol"`
}
