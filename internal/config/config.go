package config

import (
	"log"

	env "github.com/Netflix/go-env"
	_ "github.com/go-sql-driver/mysql"
)

// Config for the system
type Config struct {
	DbHost     string `env:"DB_HOST"`
	DbPort     int    `env:"DB_PORT"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASS"`
	DbName     string `env:"DB_NAME"`
}

// NewConfig operational config
func NewConfig() *Config {
	c := &Config{}
	c.DbHost = "mysql"
	c.DbPort = 3306
	c.DbUser = "root"
	c.DbPassword = ""
	c.DbName = "whatvideo"
	return c
}

// Env reads config from environment
func (c *Config) Env() *Config {
	_, err := env.UnmarshalFromEnviron(c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
