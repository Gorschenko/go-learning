package configs

type CreateUsersScriptsConfig struct {
	Count int `json:"count"`
}

type ScriptsConfig struct {
	CreateUsers CreateUsersScriptsConfig `json:"create_users"`
}
