package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	DBDatabase string `koanf:"DB_DATABASE" validate:"required"`
	DBUsername string `koanf:"DB_USERNAME" validate:"required"`
	DBPassword string `koanf:"DB_PASSWORD" validate:"required"`
	DBHost     string `koanf:"DB_HOST" validate:"required,hostname"`
	DBPort     string `koanf:"DB_PORT" validate:"required,number"`
	RedisHost  string `koanf:"REDIS_HOST" validate:"required,hostname"`
	RedisPort  int    `koanf:"REDIS_PORT" validate:"required,number"`
}

var validate *validator.Validate
var Config AppConfig

func init() {
	logrus.Info("initing config")
	k := koanf.New(".")
	validate = validator.New(validator.WithRequiredStructEnabled())
	k.Load(env.Provider("", "", func(s string) string { return s }), nil)
	k.Unmarshal("", &Config)
	err := validate.Struct(Config)

	if err != nil {
		logrus.Fatal(err)
	}
}
