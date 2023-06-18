//go:build integration

package frete_rapido_test

import (
	"github.com/ImPedro29/fr-quotation/constants"
	frete_rapido "github.com/ImPedro29/fr-quotation/lib/frete-rapido"
	"github.com/ImPedro29/fr-quotation/models"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"os"
	"testing"
)

var instance *frete_rapido.FreteRapido

func TestMain(m *testing.M) {
	l, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(l)

	instance = frete_rapido.NewFreteRapido(
		"https://sp.freterapido.com",
		"1d52a9b6b78cf07b08586152459a5c90",
		"25438296000158",
		"5AKVkHqCn",
		29161376,
		constants.DefaultHTTPTimeout,
	)

	os.Exit(m.Run())
}

func TestFreteRapido_Quote(t *testing.T) {
	request := models.QuoteRequest{
		Recipient: models.QuoteRecipient{
			Address: models.QuoteAddress{
				Zipcode: "1311000",
			},
		},
		Volumes: []models.QuoteVolume{
			{
				Category:      7,
				Amount:        1,
				UnitaryWeight: 5,
				Price:         349,
				Height:        0.2,
				Width:         0.2,
				Length:        0.2,
			},
		},
	}

	res, err := instance.Quote(request)
	require.NoError(t, err)
	require.NotEmpty(t, res)
}
