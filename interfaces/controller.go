package interfaces

import (
	"github.com/ImPedro29/fr-quotation/models"
	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	Create(ctx *fiber.Ctx) error
	Custom(route models.CustomRoute) func(ctx *fiber.Ctx) error
}
