package utils

import "fmt"

// ErrorResponse error response
type ErrorResponse struct {
	StatusCode int    `json:"-" xml:"-"`
	Request    string `json:"request" xml:"request"`
	ErrorMsg   string `json:"error" xml:"error"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("response error statusCode: %d, message: %s, request: %s", e.StatusCode, e.ErrorMsg, e.Request)
}

// GetErrorMsg get ErrorMsg from ErrorResponse
func (e *ErrorResponse) GetErrorMsg() string {
	return e.ErrorMsg
}

// GetRequest get Request from ErrorResponse
func (e *ErrorResponse) GetRequest() string {
	return e.Request
}

// GetStatusCode get StatusCode from ErrorResponse
func (e *ErrorResponse) GetStatusCode() int {
	return e.StatusCode
}
