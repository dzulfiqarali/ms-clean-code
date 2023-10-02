package shared

import (
	"net/http"
	"strconv"
)

type Error interface {
	GetError() error
	GetHTTPCode() int
	GetMessage() string
	GetCaseCode() string
}

type Base struct {
	Error           string      `json:"-"`
	StatusCode      int         `json:"-"`
	ResponseCode    string      `json:"responseCode"`
	ResponseMessage string      `json:"responseMessage"`
	AdditionalInfo  interface{} `json:"additionalInfo,omitempty"`
}

func Failure() *Base {
	return &Base{
		Error:           http.StatusText(http.StatusBadRequest),
		StatusCode:      http.StatusBadRequest,
		ResponseCode:    http.StatusText(http.StatusBadRequest),
		ResponseMessage: http.StatusText(http.StatusBadRequest),
	}
}

func Successful() *Base {
	return &Base{
		Error:           http.StatusText(http.StatusBadRequest),
		StatusCode:      http.StatusBadRequest,
		ResponseCode:    http.StatusText(http.StatusBadRequest),
		ResponseMessage: http.StatusText(http.StatusBadRequest),
	}
}
func SetStatusCode(statusCode int) func(b *Base) {
	return func(b *Base) {
		httpCode := strconv.Itoa(statusCode)
		b.ResponseCode = httpCode + ServiceCode + "00"
	}
}

func SetMessage(message string) func(b *Base) {
	return func(b *Base) {
		b.ResponseMessage = message
	}
}

func SetAdditonalInfo(data interface{}) func(b *Base) {
	return func(b *Base) {
		b.AdditionalInfo = data
	}
}

func CustomError(e Error) func(b *Base) {
	return func(b *Base) {
		b.StatusCode = e.GetHTTPCode()
		httpCode := strconv.Itoa(e.GetHTTPCode())
		b.Error = e.GetError().Error()
		b.ResponseCode = httpCode + ServiceCode + e.GetCaseCode()
		b.ResponseMessage = e.GetMessage()
	}
}
