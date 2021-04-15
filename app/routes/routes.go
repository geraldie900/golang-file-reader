/*
=====================================
;	Author : Geraldie Tanu Saputra
;	Email  : geraldie.saputra@soluix.ai
;	Date   : 15-04-2021
=====================================
*/

package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	Functions "golang-file-reader/app/functions"
	Utils "golang-file-reader/utils"
)

// InitMain function fiber
func InitMain() {
	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	app.Listen(fmt.Sprintf(":%v", Utils.PORT))
}

func setupRoutes(app *fiber.App) {
	app.Get("/file-reader", func(c *fiber.Ctx) error {
		return c.JSON(Functions.FileReader(c.Query("filename")))
	})
}
