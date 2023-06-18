package frete_rapido

import "time"

type FreteRapido struct {
	URL          string
	timeout      time.Duration
	token        string
	identity     string
	platformCode string
	cep          int64
}

type quoteRequest struct {
	Shipper        shipper      `json:"shipper"`
	Recipient      recipient    `json:"recipient"`
	Dispatchers    []dispatcher `json:"dispatchers"`
	SimulationType []int        `json:"simulation_type"`
}

type shipper struct {
	RegisteredNumber string `json:"registered_number"`
	Token            string `json:"token"`
	PlatformCode     string `json:"platform_code"`
}

type recipient struct {
	Type    int    `json:"type"`
	Country string `json:"country"`
	Zipcode int64  `json:"zipcode"`
}

type dispatcher struct {
	RegisteredNumber string   `json:"registered_number"`
	Zipcode          int64    `json:"zipcode"`
	Volumes          []volume `json:"volumes"`
}

type volume struct {
	Amount        int64   `json:"amount"`
	Category      string  `json:"category"`
	Height        float64 `json:"height"`
	Width         float64 `json:"width"`
	Length        float64 `json:"length"`
	UnitaryPrice  float64 `json:"unitary_price"`
	UnitaryWeight float64 `json:"unitary_weight"`
}

type quoteResponse struct {
	Dispatchers []responseDispatcher `json:"dispatchers"`
}

type responseDispatcher struct {
	Id                         string  `json:"id"`
	RequestId                  string  `json:"request_id"`
	RegisteredNumberShipper    string  `json:"registered_number_shipper"`
	RegisteredNumberDispatcher string  `json:"registered_number_dispatcher"`
	ZipcodeOrigin              int     `json:"zipcode_origin"`
	Offers                     []offer `json:"offers"`
}

type carrier struct {
	Name             string `json:"name"`
	RegisteredNumber string `json:"registered_number"`
	StateInscription string `json:"state_inscription"`
	Logo             string `json:"logo"`
	Reference        int    `json:"reference"`
	CompanyName      string `json:"company_name"`
}

type deliveryTime struct {
	Days          int64  `json:"days"`
	Hours         int64  `json:"hours,omitempty"`
	Minutes       int64  `json:"minutes,omitempty"`
	EstimatedDate string `json:"estimated_date"`
}

type weights struct {
	Real  int     `json:"real"`
	Used  int     `json:"used"`
	Cubed float64 `json:"cubed,omitempty"`
}

type originalDeliveryTime struct {
	Days          int    `json:"days"`
	Hours         int    `json:"hours,omitempty"`
	Minutes       int    `json:"minutes,omitempty"`
	EstimatedDate string `json:"estimated_date"`
}

type offer struct {
	Offer          int          `json:"offer"`
	TableReference string       `json:"table_reference"`
	SimulationType int          `json:"simulation_type"`
	Carrier        carrier      `json:"carrier"`
	Service        string       `json:"service"`
	DeliveryTime   deliveryTime `json:"delivery_time"`
	Expiration     time.Time    `json:"expiration"`
	CostPrice      float64      `json:"cost_price"`
	FinalPrice     float64      `json:"final_price"`
	Weights        weights      `json:"weights"`
	Correios       struct {
	} `json:"correios,omitempty"`
	OriginalDeliveryTime originalDeliveryTime `json:"original_delivery_time"`
	ServiceCode          string               `json:"service_code,omitempty"`
	ServiceDescription   string               `json:"service_description,omitempty"`
}
