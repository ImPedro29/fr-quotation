package interfaces

import "github.com/ImPedro29/fr-quotation/models"

type QuoteRepository Repository[models.Quote]
type MetricsRepository Repository[models.QuotationMetrics]

type Repository[obj any] interface {
	// CreateMany create a list of objects and returns an error
	CreateMany([]obj) error
}
