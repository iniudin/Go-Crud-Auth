package config

import (
	"log"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type Config struct {
	Port         string `env:"PORT,default=8000"`
	Database     DatabaseConfig
	JWTSecretKey string `env:"JWT_SECRET_KEY,required"`
}

type DatabaseConfig struct {
	Host     string `env:"DATABASE_HOST,default=localhost"`
	Port     string `env:"DATABASE_PORT,default=5432"`
	Username string `env:"DATABASE_USERNAME,required"`
	Password string `env:"DATABASE_PASSWORD,required"`
	Name     string `env:"DATABASE_NAME,required"`
}

func NewConfig() (*Config, error) {
	var config Config
	if err := godotenv.Load(); err != nil {
		log.Println(errors.Wrap(err, "[NewConfig] error reading .env file, defaulting to OS environment variables"))
	}

	if err := envdecode.Decode(&config); err != nil {
		return nil, errors.Wrap(err, "[NewConfig] error decoding env")
	}

	return &config, nil
}
