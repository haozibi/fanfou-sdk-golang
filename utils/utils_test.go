package utils

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowInfomation(t *testing.T) {

	ShowInfomation("Test", "abc", "haozibi")
}

func TestIsDebug(t *testing.T) {

	v := os.Getenv(debugFlag)
	defer os.Setenv(debugFlag, v)

	os.Setenv(debugFlag, "true")
	assert.Equal(t, IsDebug(), true)

	os.Setenv(debugFlag, "")
	assert.Equal(t, IsDebug(), false)

	os.Setenv(debugFlag, "on")
	assert.Equal(t, IsDebug(), false)

	os.Setenv(debugFlag, "abc")
	assert.Equal(t, IsDebug(), false)

}

func TestError(t *testing.T) {

	var f = func() error {
		return &ErrorResponse{
			StatusCode: 500,
			Request:    "/request",
			ErrorMsg:   "error_msg",
		}
	}

	err := f()

	var e *ErrorResponse
	if errors.As(err, &e) {

		assert.Equal(t, 500, e.GetStatusCode())
		assert.Equal(t, "error_msg", e.GetErrorMsg())
		assert.Equal(t, "/request", e.GetRequest())
	}

	t.Log(err)
}
