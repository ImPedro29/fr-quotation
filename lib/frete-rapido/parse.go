package frete_rapido

import (
	"fmt"
	"github.com/ImPedro29/fr-quotation/models"
	"strconv"
)

func (f *FreteRapido) parseRequest(request models.QuoteRequest) (*quoteRequest, error) {
	zipCode, err := strconv.ParseInt(request.Recipient.Address.Zipcode, 10, 64)
	if err != nil {
		return nil, fmt.Errorf(errInvalidZipCode, err)
	}

	var volumes []volume
	for _, vol := range request.Volumes {
		volumes = append(volumes, volume{
			Amount:        vol.Amount,
			Category:      fmt.Sprintf("%d", vol.Category),
			Height:        vol.Height,
			Width:         vol.Width,
			Length:        vol.Length,
			UnitaryPrice:  vol.Price,
			UnitaryWeight: vol.UnitaryWeight,
		})
	}

	internalRequest := quoteRequest{
		Shipper: shipper{
			RegisteredNumber: f.identity,
			Token:            f.token,
			PlatformCode:     f.platformCode,
		},
		Recipient: recipient{
			Type:    0,
			Country: Brazil,
			Zipcode: zipCode,
		},
		Dispatchers: []dispatcher{
			{
				RegisteredNumber: f.identity,
				Zipcode:          f.cep,
				Volumes:          volumes,
			},
		},
		SimulationType: []int{0},
	}

	return &internalRequest, nil
}

func (f *FreteRapido) parseResponse(request quoteResponse) (*models.QuoteResponse, error) {
	var carriers []models.Carrier

	if len(request.Dispatchers) < 1 {
		return nil, fmt.Errorf(errMissingDispatcher)
	}

	// always have 1 dispatcher for quotation case
	for _, aOffer := range request.Dispatchers[0].Offers {
		price := aOffer.FinalPrice
		if price == 0 {
			price = aOffer.CostPrice
		}

		carriers = append(carriers, models.Carrier{
			Name:     aOffer.Carrier.Name,
			Service:  aOffer.Service,
			Deadline: aOffer.DeliveryTime.Days,
			Price:    price,
		})
	}

	return &models.QuoteResponse{
		Carrier: carriers,
	}, nil
}
