package main

import (
	"github.com/ImPedro29/fr-quotation/server/routes"
	"github.com/ImPedro29/fr-quotation/utils"
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/zap"
	"log"
)

func main() {
	utils.LoadEnvs()

	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	logger, err := zap.NewDevelopment()
	if utils.Env.DebugMode {
		logger, err = zap.NewProduction()
		app.Use(fiberLogger.New())
	}
	if err != nil {
		log.Fatalln(err)
	}

	zap.ReplaceGlobals(logger)

	routes.Load(app)

	port := ":" + utils.Env.Port
	zap.L().Info("starting server...", zap.String("port", port))
	if err := app.Listen(port); err != nil {
		zap.L().Fatal("failed to start server", zap.Error(err))
	}
}
