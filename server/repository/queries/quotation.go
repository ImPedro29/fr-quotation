package queries

const (
	CreateQuotation     = "INSERT INTO quotations (carrier, price, days, service) VALUES (:carrier, :price, :days, :service)"
	GetQuotationMetrics = `SELECT carrier, COUNT(id) AS quantity, SUM(price) AS total_price, AVG(price) AS average, MIN(price) AS cheaper, MAX(price) AS expensive FROM (SELECT * FROM quotations ORDER BY created_at DESC %s) as data GROUP BY carrier`
)
