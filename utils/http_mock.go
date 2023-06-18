package utils

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"reflect"
)

type ctrl struct {
	statusCode int
	response   interface{}
}

// HttpMock to handle mock service to support one route
func HttpMock(requestPath string, statusCode int, response interface{}) *httptest.Server {
	c := &ctrl{statusCode, response}

	handler := http.NewServeMux()
	handler.HandleFunc(requestPath, c.mockHandler)

	return httptest.NewServer(handler)
}

func (c *ctrl) mockHandler(w http.ResponseWriter, _ *http.Request) {
	resp := []byte("{}")

	if c.response != nil {
		respType := reflect.TypeOf(c.response)
		switch respType.Kind() {
		case reflect.String:
			resp = []byte(c.response.(string))
		case reflect.Struct, reflect.Ptr:
			resp, _ = json.Marshal(c.response)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c.statusCode)
	if _, err := w.Write(resp); err != nil {
		zap.L().Warn("failed to write on route", zap.Error(err))
	}
}
