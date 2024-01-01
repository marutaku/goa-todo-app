package models

import (
	"fmt"
	"os"
)

type config struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func newPostgresConfig() *config {
	return &config{
		host:     os.Getenv("POSTGRES_HOST"),
		port:     os.Getenv("POSTGRES_PORT"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
		dbname:   os.Getenv("POSTGRES_DB"),
	}
}

func (c *config) dsn() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.host,
		c.port,
		c.user,
		c.dbname,
		c.password,
	)
}
