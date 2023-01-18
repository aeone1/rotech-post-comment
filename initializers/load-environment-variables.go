package initializers

import (
	"log"
	
	"github.com/joho/godotenv"
)

// Load Environment Variables
func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file "+err.Error())
	}
}
