package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/amber"
)
func Create() *fiber.App{
	engine := amber.New("./views",".amber")
	engine.Reload(true)
	engine.Debug(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	return app
}