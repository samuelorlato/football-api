package router

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samuelorlato/football-api/internal/integration/ports"
	"github.com/samuelorlato/football-api/pkg/errs"
)

func badRequest(msg string) error {
	return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"erro": msg})
}

func validationErrors(validator ports.Validator, err error, obj interface{}) error {
	return echo.NewHTTPError(http.StatusBadRequest, validator.GetErrors(err, obj))
}

func handleAppError(err error) error {
	var customErr *errs.Error
	if errors.As(err, &customErr) {
		return echo.NewHTTPError(customErr.Code, echo.Map{"erro": customErr.Message})
	}

	unexpected := errs.NewInternalServerError()
	return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{"erro": unexpected.Message})
}
