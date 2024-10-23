package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Host              string
	Port              string
	IdentityClientURI string
}

func NewConfig() *Config {
	_ = godotenv.Load("./deployment/.env")

	return &Config{
		Host:              os.Getenv("SELF_HOST"),
		Port:              os.Getenv("SELF_PORT"),
		IdentityClientURI: os.Getenv("IDENTITY_CLIENT_URI"),
	}
}
