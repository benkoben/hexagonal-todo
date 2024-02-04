package config

import (
	"errors"

	"github.com/caarlos0/env/v6"
)

// Server defaults
const (
    DefaultServerAddr = "localhost"
    DefaultServerPort = "8080"
)

// Database defaults
const (
    DefaultDatabaseHost = "localhost"
    DefaultDatabasePort = 5432
)

type Configuration struct {
   Database Database 
   Server Server
}

type Database struct {
    Host string `env:"DB_HOST"`
    Port uint16 `env:"DB_PORT"`
    User string `env:"DB_USERNAME"`
    Password string `env:"DB_PASSWORD"`
    Database string `env:"DB_DATABASE"`
}


type Server struct {
    Addr string `env:"SERVER_ADDR"`
    Port string `env:"SERVER_PORT"`
}

func NewConfiguration() (*Configuration, error) {
    cfg := &Configuration{
        Database: Database{
            Host: DefaultDatabaseHost,
            Port: DefaultDatabasePort,
        },
        Server: Server{
            Addr: DefaultServerAddr,
            Port: DefaultServerPort,
        },
    }
    if err := env.Parse(cfg); err != nil {
        return nil, errors.New("failed to parse configuration from environment")
    }
    return cfg, nil 
}


