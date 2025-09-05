package utils

import (
	"bookshelf/internal/entities"
	"log"

	"github.com/spf13/viper"
)

func LoadEnv() {
	_ = viper.New()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Printf("Error reading config file: %v", err)
		}
	}
}

func Env() entities.Env {
	return entities.Env{
		DatabaseURL: viper.GetString("DATABASE_URL"),
	}
}
