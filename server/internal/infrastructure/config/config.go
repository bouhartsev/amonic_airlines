package config

type Config struct {
	DatabaseConnection string `yaml:"db"`
	TokenKey           string `yaml:"token_key" env-default:"qwerty123"`
	AppPort            string `yaml:"port" env-default:":3000"`
}
