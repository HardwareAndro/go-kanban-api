package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(envName string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	return os.Getenv(envName)
}
