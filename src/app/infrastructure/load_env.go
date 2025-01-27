package infrastructure

import (
	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load environment variables
	fileName := ".env"
	err := godotenv.Load(fileName)
	if err != nil {
		panic(err)
	}
}
