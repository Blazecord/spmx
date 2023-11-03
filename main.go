package main

import (
	// "github.com/ryzmae/handlers"

	"github.com/gofiber/fiber/v2"
	"os"
	"flag"
	"log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func load() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getEnv(key string) string {
	load()

	return os.Getenv(key)
}

var (
	port = flag.String("port", getEnv("PORT"), "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	flag.Parse()

	app := fiber.New(fiber.Config{
		Prefork: *prod,
	})

	app.Use(recover.New())
	app.Use(logger.New())

	log.Fatal(app.Listen(*port))
}