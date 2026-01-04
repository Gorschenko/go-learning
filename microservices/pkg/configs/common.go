package configs

type ServicesConfig struct {
	Auth    ServiceConfig `json:"auth"`
	Devices ServiceConfig `json:"devices"`
}

type ServiceConfig struct {
	Port     int    `json:"port"`
	Host     string `json:"host"`
	Protocol string `json:"protocol"`
}
