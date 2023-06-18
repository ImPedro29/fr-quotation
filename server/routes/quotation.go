package routes

import (
	"github.com/ImPedro29/fr-quotation/constants"
	"github.com/ImPedro29/fr-quotation/server/controllers"
	"github.com/gofiber/fiber/v2"
)

func Quotation(app *fiber.App) {
	group := app.Group(string(constants.QuotationRoute))
	controller := controllers.NewQuotation()

	group.Post("", controller.Create)
	group.Get(string(constants.CustomMetricsRoute), controller.Custom(constants.CustomMetricsRoute))
}
