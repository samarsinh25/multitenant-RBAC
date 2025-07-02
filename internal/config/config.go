package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	RedisAddr  string
	RedisPwd   string
	Port       string
	Env        string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		RedisAddr:  os.Getenv("REDIS_ADDR"),
		RedisPwd:   os.Getenv("REDIS_PASSWORD"),
		Port:       os.Getenv("PORT"),
		Env:        os.Getenv("ENV"),
	}
}
