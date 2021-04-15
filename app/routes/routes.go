package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	Functions "golang-file-reader/app/functions"
)

// InitMain function fiber
func InitMain() {
	app := fiber.New()

	app.Use(logger.New())

	SetupRoutes(app)

	app.Listen(":3001")
}

func SetupRoutes(app *fiber.App) {
	app.Get("/file-upload", func(c *fiber.Ctx) error {
		return c.JSON(Functions.FileReader(c.Query("filename")))
	})
}
