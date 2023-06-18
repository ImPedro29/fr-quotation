package controllers

import (
	"github.com/ImPedro29/fr-quotation/services"
	"go.uber.org/zap"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	l, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(l)

	services.ActivateMock()

	os.Exit(m.Run())
}
