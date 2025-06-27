package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ENV struct {
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	DBPort    string
	DBSSLMode string
}

var Env ENV

func LoadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	Env = ENV{
		DBHost:    os.Getenv("DB_HOST"),
		DBUser:    os.Getenv("DB_USER"),
		DBPass:    os.Getenv("DB_PASS"),
		DBName:    os.Getenv("DB_NAME"),
		DBPort:    os.Getenv("DB_PORT"),
		DBSSLMode: os.Getenv("DB_SSLMODE"),
	}

	log.Println("ENV loading successfully!")

}
