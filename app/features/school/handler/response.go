package handler

import (
	"net/http"

	"github.com/education-hub/BE/errorr"
	"github.com/labstack/echo/v4"
)

const (
	URLFRONTEND                    = "http://localhost:8000?status=success"
	URLFRONTENDUPDATE              = "http://localhost:8000?status=success"
	URLFRONTENDFAILCREATEDGMEET    = "http://localhost:8000?status=failed"
	URLFRONTENDSUCCESSCREATEDGMEET = "http://localhost:8000?status=success"
)

type (
	WebResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    any    `json:"data,omitempty"`
	}
)

func CreateWebResponse(code int, message string, data any) any {
	return WebResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func CreateErrorResponse(err error, c echo.Context) error {
	if err, ok := err.(errorr.BadRequest); ok {
		return c.JSON(http.StatusBadRequest, CreateWebResponse(http.StatusBadRequest, err.Error(), nil))
	}
	return c.JSON(http.StatusInternalServerError, CreateWebResponse(http.StatusInternalServerError, err.Error(), nil))
}
