package configs

type JWTTokenConfig struct {
	Secret   string `json:"secret"`
	TTLHours int    `json:"ttl_hours"`
}

type JWTSecurityConfig struct {
	Access JWTTokenConfig `json:"access"`
}

type SecurityConfig struct {
	JWT JWTSecurityConfig `json:"jwt"`
}
