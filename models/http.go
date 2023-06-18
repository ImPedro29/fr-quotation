package models

import (
	"context"
	"encoding/json"
	"net/http"
)

type HTTPRequest struct {
	Ctx      context.Context
	Method   string
	URL      string
	Body     interface{}
	Response interface{}
	Headers  http.Header
}

type HTTPResponse struct {
	Data    interface{} `json:"data"`
	PerPage uint64      `json:"perPage,omitempty"`
	Page    uint64      `json:"page,omitempty"`
	Total   uint64      `json:"total,omitempty"`
}

type HTTPErrorResponse struct {
	Error        json.RawMessage `json:"error,omitempty"`
	ErrorMessage string          `json:"errorMessage,omitempty"`
	Message      string          `json:"message"`
}
