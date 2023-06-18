package models

type Quote struct {
	ID      uint64  `json:"id" db:"id"`
	Carrier string  `json:"carrier" db:"carrier"`
	Price   float64 `json:"price" db:"price"`
	Days    int64   `json:"days" db:"days"`
	Service string  `json:"service" db:"service"`
}

type QuoteRequest struct {
	Recipient QuoteRecipient `json:"recipient" validate:"required"`
	Volumes   []QuoteVolume  `json:"volumes" validate:"required,min=1"`
}

type QuoteRecipient struct {
	Address QuoteAddress `json:"address"`
}

type QuoteVolume struct {
	Category      int64   `json:"category" validate:"required"`
	Amount        int64   `json:"amount" validate:"required"`
	UnitaryWeight float64 `json:"unitary_weight" validate:"required"`
	Price         float64 `json:"price" validate:"required"`
	Sku           string  `json:"sku"`
	Height        float64 `json:"height" validate:"required"`
	Width         float64 `json:"width" validate:"required"`
	Length        float64 `json:"length" validate:"required"`
}

type QuoteAddress struct {
	Zipcode string `json:"zipcode" validate:"required,len=8"`
}

type QuoteResponse struct {
	Carrier []Carrier `json:"carrier"`
}

type Carrier struct {
	Name     string  `json:"name"`
	Service  string  `json:"service"`
	Deadline int64   `json:"deadline"`
	Price    float64 `json:"price"`
}

type QuotationMetrics struct {
	Quantity   int64   `db:"quantity" json:"quantity,omitempty"`
	TotalPrice float64 `db:"total_price" json:"totalPrice,omitempty"`
	Average    float64 `db:"average" json:"average,omitempty"`
	Cheaper    float64 `db:"cheaper" json:"cheaper,omitempty"`
	Expensive  float64 `db:"expensive" json:"expensive,omitempty"`
	Carrier    string  `db:"carrier" json:"carrier"`
}
