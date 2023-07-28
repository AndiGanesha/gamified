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
}

type Database struct {
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	Password string `env:"DB_PASSWORD"`
}

type Server struct {
	HTTP string `env:"SERVER_HTTP"`
}
