package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/config"
)

func InstallationCode(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"installation_code": config.HashedID})
}
