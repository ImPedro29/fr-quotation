package controllers

import (
	"github.com/ImPedro29/fr-quotation/constants"
	"github.com/ImPedro29/fr-quotation/interfaces"
	frete_rapido "github.com/ImPedro29/fr-quotation/lib/frete-rapido"
	"github.com/ImPedro29/fr-quotation/models"
	"github.com/ImPedro29/fr-quotation/server/repository"
	"github.com/ImPedro29/fr-quotation/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type quotation struct {
	repo interfaces.QuoteRepository
	lib  *frete_rapido.FreteRapido
}

func NewQuotation() interfaces.Controller {
	return &quotation{
		repo: repository.NewQuotation(),
		lib: frete_rapido.NewFreteRapido(
			utils.Env.FreteRapidoApi,
			utils.Env.FreteRapidoToken,
			utils.Env.FreteRapidoIdentity,
			utils.Env.FreteRapidoPlatformCode,
			utils.Env.FreteRapidoCep,
			constants.DefaultHTTPTimeout,
		),
	}
}

func (q *quotation) Create(ctx *fiber.Ctx) error {
	var data models.QuoteRequest
	if err := ctx.BodyParser(&data); err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "failed to parse body")
	}

	if err := validator.New().Struct(data); err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "invalid data")
	}

	res, err := q.lib.Quote(data)
	if err != nil {
		zap.L().Error("failed to quote", zap.Error(err))
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "failed to quote, try again")
	}

	var quotes []models.Quote
	for _, carrier := range res.Carrier {
		quotes = append(quotes, models.Quote{
			Carrier: carrier.Name,
			Price:   carrier.Price,
			Days:    carrier.Deadline,
			Service: carrier.Service,
		})
	}

	if err := q.repo.CreateMany(quotes); err != nil {
		return err
	}

	return utils.HTTPSuccess(ctx, quotes)
}

func (q *quotation) Custom(route models.CustomRoute) func(ctx *fiber.Ctx) error {
	switch route {
	case constants.CustomMetricsRoute:
		return q.Metrics
	}

	return nil
}

func (q *quotation) Metrics(ctx *fiber.Ctx) error {
	lastQuotes, err := strconv.ParseUint(ctx.Query("last_quotes", "0"), 10, 64)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "last_quotes query is invalid")
	}

	metrics, err := q.repo.(*repository.Quotation).Metrics(lastQuotes)
	if err != nil {
		zap.L().Error("failed to get metrics", zap.Error(err))
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "failed to get metrics")
	}

	return utils.HTTPSuccess(ctx, metrics)
}
