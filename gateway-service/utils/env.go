package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func Loadenv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
