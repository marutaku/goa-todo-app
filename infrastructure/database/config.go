package database

import (
	"fmt"
	"os"
)

type Config struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func NewPostgresConfig() *Config {
	return &Config{
		host:     os.Getenv("POSTGRES_HOST"),
		port:     os.Getenv("POSTGRES_PORT"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
		dbname:   os.Getenv("POSTGRES_DB"),
	}
}

func (c *Config) Dsn() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.user,
		c.password,
		c.host,
		c.port,
		c.dbname,
	)
}
