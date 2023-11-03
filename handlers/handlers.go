package handlers

import (
	"github.com/ryzmae/utils"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func GetHome(c *fiber.Ctx) error {
	err := godotenv.Load()

	if err != nil {
		return c.Status(500).SendString("Something went wrong")
	}

	blaze_url := os.Getenv("BLAZE_URL")

	return c.Redirect(blaze_url)
}

func GetHealth(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

func GetVersion(c *fiber.Ctx) error {
	version, _ := utils.GetVersion()

	return c.JSON(fiber.Map{"version": version})
}

