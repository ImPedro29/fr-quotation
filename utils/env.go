package utils

import (
	"fmt"
	"github.com/ImPedro29/fr-quotation/models"
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

// Env is loaded on start of code, can be imported on any location of the code
var Env models.Env

func LoadEnvs() {
	// ignore error because is not mandatory
	_ = godotenv.Load(".env")

	if err := env.Parse(&Env); err != nil {

		zap.L().Panic(fmt.Sprintf("%+v\n", err), zap.Error(err))
	}
}
