package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	FEURL      string
	APIKey     string
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
	PORT = os.Getenv("PORT")
	FEURL = os.Getenv("FE_URL")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")

	APIKey = os.Getenv("OPENAI_API_KEY")

}
