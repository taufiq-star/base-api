package app

import (
	"base-api/config"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Run() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	DB_USER := os.Getenv("DATABASE_USERNAME")
	DB_PASSWORD := os.Getenv("DATABASE_PASSWORD")
	DB_HOST := os.Getenv("DATABASE_HOST")
	DB_PORT := os.Getenv("DATABASE_PORT")
	DB_NAME := os.Getenv("DATABASE_NAME")

	config.Initialize(DB_USER, DB_PASSWORD, DB_PORT, DB_HOST, DB_NAME)
}
