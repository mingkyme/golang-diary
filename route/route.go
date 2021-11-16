package route

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(c.Context().QueryArgs())
		return c.Render("index", fiber.Map{})
	})
}
