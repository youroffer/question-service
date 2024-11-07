package config

import (
	"github.com/spf13/viper"
	"github.com/joho/godotenv"
)

type (
	Server struct {
		Addr string
	}

	Database struct {
		DBConn string
	}
)

type Config struct {
	Srv Server
	DB  Database
}

func New() (*Config, error) {
	// TODO: вынести в контейнер .env
	godotenv.Load()
	viper.AutomaticEnv()
	
	cfg := &Config{
		Srv: Server{
			Addr: viper.GetString("SERVER_ADDRESS"),
		},
		DB: Database{
			DBConn: viper.GetString("POSTGRES_CONN"),
		},
	}

	return cfg, nil
}
