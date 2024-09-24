package Helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(env_name string) string {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	data := os.Getenv(env_name)

	return data
}
