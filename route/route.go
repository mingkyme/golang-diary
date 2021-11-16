package route

import (
	"github.com/gofiber/fiber/v2"
)
func SetupRoute(app *fiber.App){
	app.Get("/",func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
}