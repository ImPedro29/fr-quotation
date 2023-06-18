package routes

import (
	"fmt"
	"github.com/ImPedro29/fr-quotation/constants"
	"github.com/gofiber/fiber/v2"
)

func Metrics(app *fiber.App) {
	app.Get(string(constants.MetricsRoute), func(ctx *fiber.Ctx) error {
		return ctx.Redirect(fmt.Sprintf("%s%s", constants.QuotationRoute, constants.CustomMetricsRoute))
	})
}
