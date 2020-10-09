package auth

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//GoDotEnvVariable function will get variables from .env file
func GoDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
