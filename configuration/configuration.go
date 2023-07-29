package configuration

import (
	env "github.com/Netflix/go-env"
)

func LoadConfiguration() (*Configuration, error) {
	config := &Configuration{}
	_, err := env.UnmarshalFromEnviron(config)
	return config, err
}

type Configuration struct {
	DB     Database
	Server Server
	Redis  Redis
	Token  Token
}

type Database struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
}

type Redis struct {
	Host       string `env:"REDIS_HOST"`
	Port       int    `env:"REDIS_PORT"`
	Password   string `env:"REDIS_PASSWORD"`
	ExpiryTime int    `env:"REDIS_DEFAULT_EXPIRY_TIME"`
}

type Token struct {
	AuthExpiry int `env:"TOKEN_AUTH_EXPIRY"`
}

type Server struct {
	HTTP string `env:"SERVER_HTTP"`
}
