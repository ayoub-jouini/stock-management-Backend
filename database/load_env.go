package database

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	err  := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
}