package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/ImPedro29/fr-quotation/constants"
	frete_rapido "github.com/ImPedro29/fr-quotation/lib/frete-rapido"
	"github.com/ImPedro29/fr-quotation/models"
	"github.com/ImPedro29/fr-quotation/server/repository/queries"
	"github.com/ImPedro29/fr-quotation/services"
	"github.com/ImPedro29/fr-quotation/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"net/http"
	"regexp"
	"testing"
)

func TestQuotation_Create(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		srv := utils.HttpMock(frete_rapido.QuoteRoute, http.StatusOK, returnFromFreteRapidoQuotation)
		utils.Env.FreteRapidoApi = srv.URL
		instance := NewQuotation()

		expectedQuote := models.Quote{
			Carrier: "CORREIOS",
			Price:   78.03,
			Days:    5,
			Service: "Normal",
		}

		services.Mock().ExpectBegin()
		services.Mock().ExpectExec(createQuote).
			WithArgs(expectedQuote.Carrier, expectedQuote.Price, expectedQuote.Days, expectedQuote.Service).
			WillReturnResult(sqlmock.NewResult(1, 1))
		services.Mock().ExpectCommit()

		app := fiber.New()
		fasthttpRequest := &fasthttp.RequestCtx{}
		fasthttpRequest.Request.SetBody([]byte(mockSendToApp))
		fasthttpRequest.Request.Header.Set("Content-Type", "application/json")

		c := app.AcquireCtx(fasthttpRequest)

		err := instance.Create(c)
		require.NoError(t, err)

		var result struct {
			Data []models.Quote `json:"data"`
		}
		err = json.Unmarshal(fasthttpRequest.Response.Body(), &result)
		require.NoError(t, err)
		require.NotEmpty(t, result)
		require.NotEmpty(t, result.Data)
		require.Len(t, result.Data, 1)
		assert.Equal(t, result.Data[0], expectedQuote)
	})

	t.Run("Invalid params", func(t *testing.T) {
		srv := utils.HttpMock(frete_rapido.QuoteRoute, http.StatusOK, returnFromFreteRapidoQuotation)
		utils.Env.FreteRapidoApi = srv.URL
		instance := NewQuotation()

		app := fiber.New()
		fasthttpRequest := &fasthttp.RequestCtx{}
		fasthttpRequest.Request.SetBody([]byte(mockSendToAppInvalid))
		fasthttpRequest.Request.Header.Set("Content-Type", "application/json")

		c := app.AcquireCtx(fasthttpRequest)

		err := instance.Create(c)
		require.NoError(t, err)

		var result models.HTTPErrorResponse
		err = json.Unmarshal(fasthttpRequest.Response.Body(), &result)
		require.NoError(t, err)
		require.NotEmpty(t, result)
		assert.Equal(t, "invalid data", result.Message)
	})
}

func TestQuotation_Metrics(t *testing.T) {
	expectedMetrics := models.QuotationMetrics{
		Quantity:   10,
		TotalPrice: 20,
		Average:    30,
		Cheaper:    40,
		Expensive:  50,
		Carrier:    "CORREIOS",
	}

	t.Run("Happy Path", func(t *testing.T) {
		instance := NewQuotation()

		rows := sqlmock.
			NewRows([]string{"quantity", "total_price", "average", "cheaper", "expensive", "carrier"}).
			AddRow(
				expectedMetrics.Quantity,
				expectedMetrics.TotalPrice,
				expectedMetrics.Average,
				expectedMetrics.Cheaper,
				expectedMetrics.Expensive,
				expectedMetrics.Carrier,
			)
		services.Mock().
			ExpectQuery(
				fmt.Sprintf(regexp.QuoteMeta(queries.GetQuotationMetrics), ""),
			).WillReturnRows(rows)

		app := fiber.New()
		fasthttpRequest := &fasthttp.RequestCtx{}
		fasthttpRequest.Request.Header.Set("Content-Type", "application/json")

		c := app.AcquireCtx(fasthttpRequest)
		err := instance.Custom(constants.CustomMetricsRoute)(c)
		require.NoError(t, err)

		var result struct {
			Data []models.QuotationMetrics `json:"data"`
		}
		err = json.Unmarshal(fasthttpRequest.Response.Body(), &result)
		require.NoError(t, err)
		require.NotEmpty(t, result)
		require.NotEmpty(t, result.Data)
		require.Len(t, result.Data, 1)
		assert.Equal(t, result.Data[0], expectedMetrics)
	})

	t.Run("Happy Path - With params", func(t *testing.T) {
		instance := NewQuotation()

		rows := sqlmock.
			NewRows([]string{"quantity", "total_price", "average", "cheaper", "expensive", "carrier"}).
			AddRow(
				expectedMetrics.Quantity,
				expectedMetrics.TotalPrice,
				expectedMetrics.Average,
				expectedMetrics.Cheaper,
				expectedMetrics.Expensive,
				expectedMetrics.Carrier,
			)
		services.Mock().ExpectQuery(
			fmt.Sprintf(regexp.QuoteMeta(queries.GetQuotationMetrics), " LIMIT 10"),
		).WillReturnRows(rows)

		app := fiber.New()
		fasthttpRequest := &fasthttp.RequestCtx{}
		fasthttpRequest.Request.Header.Set("Content-Type", "application/json")
		fasthttpRequest.Request.SetRequestURI(string(fasthttpRequest.Request.RequestURI()) + "?last_quotes=10")

		c := app.AcquireCtx(fasthttpRequest)
		err := instance.Custom(constants.CustomMetricsRoute)(c)
		require.NoError(t, err)

		var result struct {
			Data []models.QuotationMetrics `json:"data"`
		}
		err = json.Unmarshal(fasthttpRequest.Response.Body(), &result)
		require.NoError(t, err)
		require.NotEmpty(t, result)
		require.NotEmpty(t, result.Data)
		require.Len(t, result.Data, 1)
		assert.Equal(t, result.Data[0], expectedMetrics)
	})

	t.Run("Invalid param", func(t *testing.T) {
		instance := NewQuotation()

		app := fiber.New()
		fasthttpRequest := &fasthttp.RequestCtx{}
		fasthttpRequest.Request.Header.Set("Content-Type", "application/json")
		fasthttpRequest.Request.SetRequestURI(string(fasthttpRequest.Request.RequestURI()) + "?last_quotes=abc")

		c := app.AcquireCtx(fasthttpRequest)
		err := instance.Custom(constants.CustomMetricsRoute)(c)
		require.NoError(t, err)

		var result models.HTTPErrorResponse
		err = json.Unmarshal(fasthttpRequest.Response.Body(), &result)
		require.NoError(t, err)
		require.NotEmpty(t, result)
		assert.Equal(t, "last_quotes query is invalid", result.Message)
	})
}
