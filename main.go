package main

import (
	"errors"
	"log"
	"os"

	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
	"github.com/olufekosamuel/voiceout-api/routes"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	logger, _ = thoth.Init("log")
)

func init() {
	if err := godotenv.Load(); err != nil {
		logger.Log(errors.New("no .env file found"))
		log.Fatal("No .env File Found")
	}

	connectionString := os.Getenv("DATABASE_URL")

	err := mgm.SetDefaultConfig(nil, "voiceout", options.Client().ApplyURI(connectionString))
	if err != nil {
		logger.Log(errors.New("cannot set mgm config"))
		log.Fatal(err)
	}
}

func main() {

	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024,
	})
	app.Use(cors.New())
	routes.RegisterRoutes(app)
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "5000"
	}
	port = ":" + port

	log.Fatal(app.Listen(port))
}
