package routes

import "github.com/gofiber/fiber/v2"

func Load(app *fiber.App) {
	Quotation(app)
	Metrics(app)
}
