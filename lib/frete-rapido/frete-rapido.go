package frete_rapido

import (
	"context"
	"fmt"
	"github.com/ImPedro29/fr-quotation/models"
	"github.com/ImPedro29/fr-quotation/utils"
	"net/http"
	"time"
)

func NewFreteRapido(url, token, identity, platformCode string, cep int64, timeout time.Duration) *FreteRapido {
	return &FreteRapido{
		URL:          url,
		timeout:      timeout,
		token:        token,
		identity:     identity,
		platformCode: platformCode,
		cep:          cep,
	}
}

func (f *FreteRapido) Quote(req models.QuoteRequest) (*models.QuoteResponse, error) {
	ctx, c := context.WithTimeout(context.TODO(), f.timeout)
	defer c()

	url := fmt.Sprintf("%s%s", f.URL, QuoteRoute)

	request, err := f.parseRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to parse request: %w", err)
	}

	var response quoteResponse
	if err := utils.Request(models.HTTPRequest{
		Ctx:      ctx,
		Method:   http.MethodPost,
		URL:      url,
		Body:     request,
		Response: &response,
	}); err != nil {
		return nil, err
	}

	return f.parseResponse(response)
}
