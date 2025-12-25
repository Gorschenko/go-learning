package configs

import "time"

type JWTTokenConfig struct {
	Secret string        `json:"secret"`
	TTL    time.Duration `json:"ttl_hours"`
}

type JWTSecurityConfig struct {
	Access JWTTokenConfig `json:"access"`
}

type SecurityConfig struct {
	JWT JWTSecurityConfig `json:"jwt"`
}
