package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct { // структура для хранения конфигурации приложения
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct { // структура для хранения конфигурации базы данных
	Dsn string
}

type AuthConfig struct { // структура для хранения конфигурации аутентификации
	Secret string
}

func LoadConfig() *Config { // загрузка конфигурации из .env файла
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
	}
}
