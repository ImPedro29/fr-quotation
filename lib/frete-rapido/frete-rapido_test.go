package frete_rapido

import (
	"github.com/ImPedro29/fr-quotation/constants"
	"github.com/ImPedro29/fr-quotation/models"
	"github.com/ImPedro29/fr-quotation/utils"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

var instance = NewFreteRapido(
	"https://sp.freterapido.com",
	"1d52a9b6b78cf07b08586152459a5c90",
	"25438296000158",
	"5AKVkHqCn",
	29161376,
	constants.DefaultHTTPTimeout,
)

func TestFreteRapido_Quote(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		srv := utils.HttpMock(QuoteRoute, http.StatusOK, happyPathSimulate)
		instance.URL = srv.URL

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
		require.Equal(t, 1, len(res.Carrier))
		require.Equal(t, res.Carrier[0], models.Carrier{
			Name:     "CORREIOS",
			Service:  "Normal",
			Deadline: 5,
			Price:    78.03,
		})
	})

	t.Run("Bad request", func(t *testing.T) {
		srv := utils.HttpMock(QuoteRoute, http.StatusBadRequest, happyPathSimulate)
		instance.URL = srv.URL

		request := models.QuoteRequest{
			Recipient: models.QuoteRecipient{
				Address: models.QuoteAddress{
					Zipcode: "1",
				},
			},
		}

		res, err := instance.Quote(request)
		require.Error(t, err)
		require.Empty(t, res)
	})
}
