package main

import (
	"GID/controllers"
	"GID/helpers"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}
	r := controllers.SetupRouter()
	err := r.Run()
	if err != nil {
		helpers.Log.Fatal(err)
		return
	}
}
