package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onumahkalusamuel/hospital-app/pkg"
)

type ActivationRequest struct {
	InstallationCode string `json:"installation_code"`
}

func GetActivationCode(c echo.Context) error {
	// edit later to tie hosiptal name to activation code

	var activationRequest ActivationRequest
	err := c.Bind(&activationRequest)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if len(activationRequest.InstallationCode) < 10 {
		return c.JSON(http.StatusOK, echo.Map{
			"code":    3,
			"success": false,
			"message": "Invalid installation code",
		})
	}

	activationCode := pkg.GenActivationCode(activationRequest.InstallationCode)

	return c.JSON(http.StatusOK, echo.Map{
		"success":         true,
		"activation_code": activationCode,
	})
}
