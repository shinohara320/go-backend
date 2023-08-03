package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	routes "github.com/shinohara320/travel-agent/Routes"
	"github.com/shinohara320/travel-agent/database"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}
	port := os.Getenv("PORT")
	certFile := os.Getenv("CERT_FILE") // Path to your cert.pem
	keyFile := os.Getenv("KEY_FILE")   // Path to your key.pem

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	// Enable HTTPS
	err = app.ListenTLS(":"+port, certFile, keyFile)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	routes.Setup(app)
}
