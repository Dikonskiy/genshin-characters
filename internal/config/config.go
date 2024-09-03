package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	GetPort() string
	GetConnection() string
}

type config struct {
	appPort      string
	dbConnection string
}

func NewConfig(filename string) (Config, error) {
	if err := godotenv.Load(filename); err != nil {
		return nil, err
	}

	config := &config{
		appPort:      os.Getenv("APP_PORT"),
		dbConnection: os.Getenv("MYSQL_CONNECTION"),
	}

	return config, nil
}

func (c *config) GetPort() string {
	return c.appPort
}

func (c *config) GetConnection() string {
	return c.dbConnection
}
