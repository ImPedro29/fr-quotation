package repository

import (
	"context"
	"fmt"
	"github.com/ImPedro29/fr-quotation/constants"
	"github.com/ImPedro29/fr-quotation/interfaces"
	"github.com/ImPedro29/fr-quotation/models"
	"github.com/ImPedro29/fr-quotation/server/repository/queries"
	"github.com/ImPedro29/fr-quotation/services"
	"github.com/jmoiron/sqlx"
	"time"
)

type Quotation struct {
	db      *sqlx.DB
	timeout time.Duration
}

func NewQuotation() interfaces.QuoteRepository {
	return &Quotation{
		db:      services.Postgres(),
		timeout: constants.DBTimeout,
	}
}

func (q *Quotation) Metrics(limit uint64) ([]models.QuotationMetrics, error) {
	ctx, c := context.WithTimeout(context.TODO(), q.timeout)
	defer c()

	limitStr := ""
	if limit > 0 {
		limitStr = fmt.Sprintf(" LIMIT %d", limit)
	}

	query := fmt.Sprintf(queries.GetQuotationMetrics, limitStr)
	var data []models.QuotationMetrics
	err := q.db.SelectContext(ctx, &data, query)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (q *Quotation) CreateMany(m []models.Quote) error {
	ctx, c := context.WithTimeout(context.TODO(), q.timeout)
	defer c()

	tx, err := q.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = q.db.NamedExecContext(ctx, queries.CreateQuotation, m)
	if err != nil {
		return err
	}

	return tx.Commit()
}
