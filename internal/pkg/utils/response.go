package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RespondJSONError(e echo.Context, code int, err interface{}) error {
	return respondJSON(e, code, map[string]interface{}{
		"err":    err,
		"status": code,
	})
}

func RespondJSONSuccess(e echo.Context, data interface{}) error {
	return respondJSON(e, http.StatusOK, map[string]interface{}{
		"data":   data,
		"status": http.StatusOK,
	})
}

func respondJSON(e echo.Context, code int, payload interface{}) error {
	return e.JSON(code, payload)
}
