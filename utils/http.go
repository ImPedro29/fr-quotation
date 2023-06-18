package utils

import (
	"encoding/json"
	"github.com/ImPedro29/fr-quotation/constants"
	"github.com/ImPedro29/fr-quotation/models"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"io"
	"net/http"
	"reflect"
	"strings"
)

var client = &http.Client{
	Timeout: constants.DefaultHTTPTimeout,
}

// Request can be used as a generalist method to do a http request
func Request(request models.HTTPRequest) error {
	bodyParsed, err := json.Marshal(request.Body)
	if err != nil {
		return err
	}

	if request.Headers == nil {
		request.Headers = make(http.Header)
	}

	request.Headers.Set("Content-Type", "application/json")
	httpRequest, err := http.NewRequestWithContext(request.Ctx, request.Method, request.URL, strings.NewReader(string(bodyParsed)))
	if err != nil {
		return err
	}
	httpRequest.Header = request.Headers

	zap.L().Debug("performing request", zap.Any("request", request))
	res, err := client.Do(httpRequest)
	if err != nil {
		return err
	}

	contentType := res.Header.Get("Content-Type")
	if res.StatusCode >= http.StatusBadRequest {
		if strings.HasPrefix(contentType, "application/json") {
			if err := json.NewDecoder(res.Body).Decode(request.Response); err != nil {
				zap.L().Warn("failed to decode json to passed struct", zap.Error(err))
			}
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		zap.L().Error("request not returned 200", zap.Any("request", httpRequest), zap.ByteString("result", data))
		return ErrRequestStatusCode
	}

	if strings.HasPrefix(contentType, "application/json") {
		return json.NewDecoder(res.Body).Decode(request.Response)
	}

	if strings.HasPrefix(contentType, "text/plain") {
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		dataStr := `"` + string(data) + `"`
		return json.Unmarshal([]byte(dataStr), request.Response)
	}

	return nil
}

func HTTPSuccess(ctx *fiber.Ctx, data interface{}) error {
	dataReflect := reflect.ValueOf(data)
	isSlice := dataReflect.Kind() == reflect.Slice

	if data == nil || (isSlice && dataReflect.Len() < 1) {
		data = []string{}
	}

	return ctx.Status(http.StatusOK).JSON(&models.HTTPResponse{
		Data: data,
	})
}

func HTTPFail(ctx *fiber.Ctx, code int, err error, message string) error {
	errJson, _ := json.Marshal(err)

	result := &models.HTTPErrorResponse{
		Error:   errJson,
		Message: message,
	}

	if err != nil {
		result.ErrorMessage = err.Error()
	}

	return ctx.Status(code).JSON(result)
}
