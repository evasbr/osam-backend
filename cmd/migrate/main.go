package main

import (
	"log"

	"github.com/evasbr/osam-backend/app/config"
	"github.com/evasbr/osam-backend/app/model"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config.ConnectToDB()

	err = config.DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("AutoMigration failed:", err)
	}
	log.Println("Migration success")
}
